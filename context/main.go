package main

import (
	"context"
	"fmt"
	"time"
)

var slowAPIresponseTime = 180

const answerToEverything = 42
const youStupid = 21

type Response struct {
	value int
	err   error
}

func handleAPICall(timeoutMilliseconds int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMilliseconds)*time.Millisecond)
	defer cancel()
	responseCh := make(chan Response)

	go func() {
		value, err := slowAPICall(ctx)
		responseCh <- Response{value: value, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return youStupid, fmt.Errorf("API call timed out")
		case response := <-responseCh:
			return response.value, response.err
		}
	}
}

func slowAPICall(ctx context.Context) (int, error) {
	time.Sleep(time.Millisecond * time.Duration(slowAPIresponseTime))
	return answerToEverything, nil
}

func main() { fmt.Println(handleAPICall(100)) }
