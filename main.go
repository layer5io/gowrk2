package main

import (
	"github.com/layer5io/gowrk2/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	thread := 2
	duration := 5
	connection := 10
	rqps := 10
	url := "https://gmail.com:443"
	// scriptLua := "/scripts/multiple-endpoints_in_json.lua"
	result, _ := api.WRKRun(thread, duration, connection, rqps, url)
	logrus.Infof("WRK Result: %+v", result)

	fortioResult, _ := api.TransformWRKToFortio(result)
	logrus.Infof("Fortio Result: %+v", fortioResult)
}
