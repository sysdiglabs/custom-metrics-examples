package main

import (
	"github.com/quipo/statsd"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create the client
	client := statsd.NewStatsdClient("127.0.0.1:8125", "golang.")
	defer client.Close()
	// Connect to the statsd server
	err := client.CreateSocket()
	if err != nil {
		panic(err)
	}

	for {

		// Gauge
		client.Gauge("foo", rand.Int63n(100))

		// Counter
		client.Incr("bar", 1)

		// Counter with sampling
		client.IncrWithSampling("bar_sampled", 5, 0.1)

		// Timer
		client.Timing("timed", rand.Int63n(100000))

		// Wait for a random time between 500 and 1000 ms before sending the data again.
		time.Sleep(time.Duration(int(time.Millisecond) * (rand.Intn(500) + 500)))
	}

}
