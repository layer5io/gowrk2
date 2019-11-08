package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/layer5io/gowrk2/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	config := &api.GoWRK2Config{
		Thread:      2,
		Duration:    5,
		Connection:  10,
		RQPS:        10,
		URL:         "https://gmail.com:443",
		Percentiles: []float64{50, 75, 90, 99, 99.99, 99.999},
		Labels:      "",
	}
	result, _ := api.WRKRun(config)
	logrus.Infof("WRK Result: %+v", result)

	fortioResult, _ := api.TransformWRKToFortio(result, config)
	// logrus.Infof("Fortio Result: %+#v", fortioResult)
	logrus.Info("Fortio Result")
	spew.Dump(fortioResult)
}
