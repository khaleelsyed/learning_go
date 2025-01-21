package context

import (
	"context"
	"fmt"
	"time"
)

type APIResponse int

const answerToEverything APIResponse = 42
const youStupid APIResponse = 21

type APICaller interface {
	CallAPI(ctx context.Context) APIResponse
}

type SlowAPICaller struct{}
type FastAPICaller struct{}

func handleAPICall(timeoutMilliseconds int, apiCaller APICaller) (APIResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMilliseconds)*time.Millisecond)
	defer cancel()
	responseCh := make(chan APIResponse)

	go func() {
		response := apiCaller.CallAPI(ctx)
		responseCh <- response
	}()

	for {
		select {
		case <-ctx.Done():
			return youStupid, fmt.Errorf("API call timed out")
		case response := <-responseCh:
			return response, nil
		}
	}
}

func NewSlowAPICaller() SlowAPICaller {
	return SlowAPICaller{}
}

func NewFastAPICaller() FastAPICaller {
	return FastAPICaller{}
}

func (api SlowAPICaller) CallAPI(ctx context.Context) APIResponse {
	time.Sleep(time.Millisecond * 350)
	return answerToEverything
}

func (api FastAPICaller) CallAPI(ctx context.Context) APIResponse {
	time.Sleep(time.Millisecond * 98)
	return answerToEverything
}
