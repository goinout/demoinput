package main

import (
	"context"
	"fmt"
)

func stop() {
	fmt.Println("demoinput stopped")
}

// Exchange contexts with host
func Exchange(ctx context.Context) context.Context {
	go func() {
		select {
		case <-ctx.Done():
			// listen for the host cancel signal
			stop()
		}
	}()

	// no need for host to be aware of this plugin stopped
	return nil
}

// Start the input
func Start(getOutput func(string) func(map[string]interface{}) error) {
	fmt.Println("demoinput started")
	fn := getOutput("demoouput")
	if fn == nil {
		fmt.Println("demooutput can't be found")
	} else {
		demoParam := make(map[string]interface{})
		demoParam["demo"] = "hello"

		err := fn(demoParam)
		fmt.Println("err:", err)
	}
}
