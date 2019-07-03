package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	client, err := statsd.New(os.Getenv("DD_AGENT_HOST") + ":" + os.Getenv("DD_DOGSTATSD_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DSD client initialized")
	for {
		err = client.Gauge("custom.metric", 1.5, nil, 1)
		log.Printf("sending metric")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 15)
	}
}
