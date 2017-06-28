package main

import (
	"context"
	"log"
)

func stop() {
	log.Println("demoinput stopped")
}

// Exchange contexts with host
func Exchange(ctx interface{}) interface{} {
	c := ctx.(context.Context)
	go func() {
		select {
		case <-c.Done():
			// listen for the host cancel signal
			stop()
		}
	}()

	// no need for host to be aware of this plugin stopped
	return nil
}

// Start the input
func Start(getOutput func(string) func(map[string]interface{}) error) {
	log.Println("demoinput started")
	fn := getOutput("demoouput")
	if fn == nil {
		log.Println("demooutput can't be found")
	} else {
		demoParam := make(map[string]interface{})
		demoParam["demo"] = "hello"

		err := fn(demoParam)
		log.Println("err:", err)
	}
}
