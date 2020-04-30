package main

import (
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

// This isn't really necessary I don't think
const interval time.Duration = 15

// Start the things!
func initStatsdClient() *statsd.Client {
	// Host env var
	host := os.Getenv("DD_AGENT_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	// Port env var
	port := os.Getenv("DD_DOGSTATSD_PORT")
	if port == "" {
		port = "8125"
	}
	client, err := statsd.New(host + ":" + port)
	log.Println("Trying dd-agent UDP address: " + host + ":" + port)

	socket := os.Getenv("DD_DOGSTATSD_SOCKET")
	if socket != "" {
		client, err = statsd.New(statsd.UnixAddressPrefix + socket)
		log.Println("Trying dd-agent UDS address: " + statsd.UnixAddressPrefix + socket)
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

// Do the things!
func main() {
	client := initStatsdClient()
	if client != nil {
		log.Println("Dogstatsd client successfully initialized")
	}
	for {
		err := client.Gauge("datadog.custom.metric", 1.5, nil, 1)
		log.Printf("sending metric: custom.metric")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * interval)
	}
}
