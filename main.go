package main

import (
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

const interval time.Duration = 15

func initStatsdClient() *statsd.Client {
	host := os.Getenv("DD_AGENT_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("DD_DOGSTATSD_PORT")
	if port == "" {
		port = "8125"
	}
	client, err := statsd.New(host + ":" + port)
	log.Println("trying dd-agent udp addr: " + host + ":" + port)

	socket := os.Getenv("DD_DOGSTATSD_SOCKET")
	if socket != "" {
		client, err = statsd.New(statsd.UnixAddressPrefix + socket)
		log.Println("trying dd-agent uds addr: " + statsd.UnixAddressPrefix + socket)
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

func main() {
	client := initStatsdClient()
	if client != nil {
		log.Println("DSD client successfully initialized")
	}
	for {
		err := client.Gauge("custom.metric", 1.5, nil, 1)
		log.Printf("sending metric: custom.metric")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * interval)
	}
}
