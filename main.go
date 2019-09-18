package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func wrkRun(thread, duration, connection, rqps int, url string, scriptLua string) []byte {

	out, err := exec.Command("wrk", "-t"+strconv.Itoa(thread), "-d"+strconv.Itoa(duration)+"s", "-c"+strconv.Itoa(connection), "-R"+strconv.Itoa(rqps), "-s", scriptLua, url).Output()
	if err != nil {
		log.Fatal(err)
	}
	in := []byte(out)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	result, _ := json.Marshal(raw)

	return result

}

func main() {
	thread := 2
	duration := 2
	connection := 10
	rqps := 10
	url := "https://gmail.com:443"
	scriptLua := "/scripts/multiple-endpoints_in_json.lua"
	result := wrkRun(thread, duration, connection, rqps, url, scriptLua)
	fmt.Printf(string(result))

}
