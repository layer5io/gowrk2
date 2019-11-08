package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	"time"

	"fortio.org/fortio/fhttp"
	"fortio.org/fortio/periodic"
	"fortio.org/fortio/stats"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GoWRK2 struct {
	TotalRequests          int64            `json:"TotalRequests"`
	DurationInMicroseconds float64          `json:"DurationInMicroseconds"`
	Bytes                  int64            `json:"Bytes"`
	RequestsPerSec         float64          `json:"RequestsPerSec"`
	BytesTransferPerSec    float64          `json:"BytesTransferPerSec"`
	Errors                 float64          `json:"Errors"`
	MinLatency             float64          `json:"MinLatency"`
	AvgLatency             float64          `json:"AvgLatency"`
	MaxLatency             float64          `json:"MaxLatency"`
	StdDev                 float64          `json:"StdDev"`
	URL0                   string           `json:"Url_0"`
	URLRequestCount0       int              `json:"UrlRequestCount_0"`
	Percentiles            []PercentileInfo `json:"Percentiles"`
}

type PercentileInfo struct {
	Count   int64   `json:"Count"`
	Percent float64 `json:"Percent"`
	Value   float64 `json:"Value"`
}

type GoWRK2Config struct {
	Thread                  int
	DurationInSeconds, RQPS float64
	URL, Labels             string
	Percentiles             []float64
}

func WRKRun(config *GoWRK2Config) (*GoWRK2, error) {
	scriptLua := "./wrk2/scripts/multiple-endpoints_in_json.lua"
	args := []string{"-t" + strconv.Itoa(config.Thread),
		"-d" + strconv.FormatFloat(config.DurationInSeconds, 'f', -1, 64) + "s",
		"-R" + strconv.FormatFloat(config.RQPS, 'f', -1, 64),
		"-s", scriptLua, config.URL}
	logrus.Debugf("received command: wrk %v", args)
	out, err := exec.Command("wrk", args...).Output()
	if err != nil {
		err = errors.Wrapf(err, "unable to execute the requested command")
		logrus.Error(err)
		return nil, err
	}
	logrus.Debugf("Received output: %s", out)
	in := []byte(out)
	var raw *GoWRK2
	if err := json.Unmarshal(in, &raw); err != nil {
		err = errors.Wrapf(err, "unable to marshal the result")
		logrus.Error(err)
		return nil, err
	}
	return raw, nil
}

func TransformWRKToFortio(gowrk *GoWRK2, config *GoWRK2Config) (*fhttp.HTTPRunnerResults, error) {
	if gowrk != nil {
		dur, err := time.ParseDuration(fmt.Sprintf("%fus", gowrk.DurationInMicroseconds))
		if err != nil {
			err = errors.Wrapf(err, "unable to parse duration in microseconds")
			logrus.Error(err)
			return nil, err
		}
		logrus.Debugf("parsed duration string: %f to dur: %v", gowrk.DurationInMicroseconds, dur)

		result := &fhttp.HTTPRunnerResults{
			// we dont intend to support multiple URLs at the moment
			URL: gowrk.URL0,
			RunnerResults: periodic.RunnerResults{
				Labels:         config.Labels,
				RunType:        "HTTP",
				ActualDuration: dur,
				ActualQPS:      gowrk.RequestsPerSec,
				DurationHistogram: &stats.HistogramData{
					Avg:   gowrk.AvgLatency / 1000,
					Min:   gowrk.MinLatency / 1000,
					Max:   gowrk.MaxLatency / 1000,
					Count: gowrk.TotalRequests,
				},
			},
		}

		var countTrkr int64
		var windowTrkr float64
		for _, p := range gowrk.Percentiles {
			for _, pr := range config.Percentiles {
				if p.Percent == pr {
					result.DurationHistogram.Percentiles = append(result.DurationHistogram.Percentiles, stats.Percentile{
						Value:      p.Value,
						Percentile: p.Percent,
					})
				}
			}
			result.DurationHistogram.Data = append(result.DurationHistogram.Data, stats.Bucket{
				Count:   p.Count - countTrkr,
				Percent: p.Percent,
				Interval: stats.Interval{
					Start: windowTrkr,
					End:   p.Value,
				},
			})
			countTrkr = p.Count
			windowTrkr = p.Value
		}

		return result, nil
	}

	return nil, nil
}
