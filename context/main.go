package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const answerToEverything = 42

func handleAPICall(timeoutSeconds int) {
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Duration(timeoutSeconds)*time.Second, fmt.Errorf("%s", fmt.Sprintf("Over %d seconds have already elapsed", timeoutSeconds)))
	defer cancel()

	for {
		go func(ctx context.Context) {
			value, err := slowAPICall(ctx)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(value)
				return
			}
		}(ctx)
		time.Sleep(time.Millisecond * 500)
	}
}

func slowAPICall(ctx context.Context) (int, error) {
	time.Sleep(time.Second * 1)
	return answerToEverything, nil
}

func main() { handleAPICall(2) }
