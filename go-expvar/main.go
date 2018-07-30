package main

import (
	"expvar"
	"net/http"
	"time"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

// Publish the port as soon as the program starts
func init() {
	go http.ListenAndServe(":8080", nil)
}

// Custom struct that will be exported
type Load struct {
	Load1  float64
	Load5  float64
	Load15 float64
}

// Function that will be called by expvar
// to export the information from the structure
// every time the endpoint is reached
func AllLoadAvg() interface{} {
	return Load{
		Load1:  loadAvg(0),
		Load5:  loadAvg(1),
		Load15: loadAvg(2),
	}
}

// Aux function to retrieve the load average
// in GNU/Linux systems
func loadAvg(position int) float64 {
	data, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err)
	}
	values := strings.Fields(string(data))

	load, err := strconv.ParseFloat(values[position], 64)
	if err != nil {
		panic(err)
	}

	return load
}

func main() {

	var (
		numberOfSecondsRunning = expvar.NewInt("system.numberOfSeconds")
		programName            = expvar.NewString("system.programName")
		lastLoad               = expvar.NewFloat("system.lastLoad")
		numberOfLoginsPerUser  = expvar.NewMap("system.numberOfLoginsPerUser")
	)

	// The contents returned by the function will be autoexported in JSON format
	expvar.Publish("system.allLoad", expvar.Func(AllLoadAvg))

	programName.Set(os.Args[0])

	// We will increment this metrics every second
	for {
		numberOfSecondsRunning.Add(1)
		lastLoad.Set(loadAvg(0))
		numberOfLoginsPerUser.Add("foo", 2)
		numberOfLoginsPerUser.Add("bar", 1)
		time.Sleep(1 * time.Second)
	}
}
