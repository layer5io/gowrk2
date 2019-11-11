package main

import (
	"encoding/json"

	"github.com/layer5io/gowrk2/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	config := &api.GoWRK2Config{
		Thread:            1,
		DurationInSeconds: 5,
		// Connection:        10,
		RQPS:        10,
		URL:         "http://gmail.com",
		Percentiles: []float64{50, 75, 90, 99, 99.99, 99.999},
		Labels:      "",
	}
	result, _ := api.WRKRun(config)
	logrus.Infof("WRK Result: %+v", result)

	fortioResult, _ := api.TransformWRKToFortio(result, config)
	// logrus.Infof("Fortio Result: %+#v", fortioResult)
	logrus.Info("Fortio Result")
	// spew.Dump(fortioResult)
	jb, _ := json.Marshal(fortioResult)
	logrus.Infof("json: %s", jb)
}
