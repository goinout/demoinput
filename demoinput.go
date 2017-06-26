package main

import (
	"context"
	"fmt"

	"github.com/goinout/goinout"
)

func stop() {
	fmt.Println("stopped")
}

// Start the input
func Start(ctx context.Context, outputs goinout.OutputGetter) {
	go func() {
		select {
		case <-ctx.Done():
			stop()
		}
	}()

	fmt.Println("demoinput started")
	fn := outputs("demoouput")
	if fn == nil {
		fmt.Println("demooutput can't be found")
	} else {
		demoParam := make(map[string]interface{})
		demoParam["demo"] = "hello"

		data := []byte("demo data")
		err := fn(demoParam, data)
		fmt.Println("err", err)
	}
}
