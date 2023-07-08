package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"time"

	"fortio.org/fortio/fhttp"
	"fortio.org/fortio/periodic"
	"fortio.org/fortio/stats"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GoWRK2 struct {
	RequestedDuration string    `json:"RequestedDuration,omitempty"`
	RequestedQPS      string    `json:"RequestedQPS,omitempty"`
	StartTime         time.Time `json:"StartTime,omitempty"`

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
	wrkLoc := "./wrk2/wrk"
	wrkLocFromEnv := os.Getenv("WRK_LOCATION")
	if wrkLocFromEnv != "" {
		wrkLoc = wrkLocFromEnv
	}
	rURLI, _ := url.Parse(config.URL)
	if !rURLI.IsAbs() {
		err := fmt.Errorf("given URL (%s) is not a valid URL", config.URL)
		logrus.Error(err)
		return nil, err
	}

	if rURLI.Port() == "" {
		if rURLI.Scheme == "https" {
			rURLI.Host += ":443"
		} else {
			rURLI.Host += ":80"
		}
	}
	dur := strconv.FormatFloat(config.DurationInSeconds, 'f', -1, 64)
	scriptLua := "./wrk2/scripts/multiple-endpoints_in_json.lua"
	args := []string{"-t" + strconv.Itoa(config.Thread),
		"-d" + dur + "s",
		"-R" + strconv.FormatFloat(config.RQPS, 'f', -1, 64),
		"-s", scriptLua, rURLI.String()}
	logrus.Debugf("received command: wrk %v", args)

	var startTime time.Time
	go func() {
		startTime = time.Now()
	}()
	out, err := exec.Command(wrkLoc, args...).Output()
	if err != nil {
		err = errors.Wrapf(err, "unable to execute the requested command")
		logrus.Error(err)
		return nil, err
	}

	retryCount := 0

RETRY:
	logrus.Debugf("Received output: %s", out)
	var raw *GoWRK2
	if err := json.Unmarshal(out, &raw); err != nil {
		retryCount++
		if retryCount <= 1 {
			logrus.Debugf("starting with bytes: %s", out)
			in := string(out)
			ind := strings.Index(in, `,`)
			logrus.Debugf("index of the first new line: %d", ind)
			if ind > -1 && ind+1 < len(in) {
				logrus.Debug("there was an unmarshal error, retrying with a tiny change")
				in = in[ind+1:]
				out = []byte(in)
				logrus.Debugf("new output: %s", out)
				goto RETRY
			}
		} else {
			err = errors.Wrapf(err, "unable to marshal the result")
			logrus.Error(err)
			return nil, err
		}
	}

	raw.StartTime = startTime
	raw.RequestedDuration = dur + "s"
	raw.RequestedQPS = fmt.Sprintf("%f", config.RQPS)

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
			RunnerResults: periodic.RunnerResults{
				StartTime:         gowrk.StartTime,
				RequestedQPS:      gowrk.RequestedQPS,
				RequestedDuration: gowrk.RequestedDuration,
				NumThreads:        config.Thread,
				Labels:            config.Labels,
				RunType:           "HTTP",
				ActualDuration:    dur,
				ActualQPS:         gowrk.RequestsPerSec,
				DurationHistogram: &stats.HistogramData{
					Avg:    gowrk.AvgLatency / 1000000,
					Min:    gowrk.MinLatency / 1000000,
					Max:    gowrk.MaxLatency / 1000000,
					Count:  gowrk.TotalRequests,
					StdDev: gowrk.StdDev / 1000000,
				},
			},
		}

		result.HTTPOptions = fhttp.HTTPOptions{
			URL: gowrk.URL0,
		}

		// var countTrkr int64
		var windowTrkr float64
		for _, p := range gowrk.Percentiles {
			for _, pr := range config.Percentiles {
				if p.Percent == pr {
					result.DurationHistogram.Percentiles = append(result.DurationHistogram.Percentiles, stats.Percentile{
						Value:      p.Value / 1000,
						Percentile: p.Percent,
					})
				}
			}
			result.DurationHistogram.Data = append(result.DurationHistogram.Data, stats.Bucket{
				// Count:   p.Count - countTrkr,
				Count:   p.Count,
				Percent: p.Percent,
				Interval: stats.Interval{
					Start: windowTrkr,
					End:   p.Value / 1000,
				},
			})
			// countTrkr = p.Count
			windowTrkr = p.Value / 1000
		}

		return result, nil
	}

	return nil, nil
}
