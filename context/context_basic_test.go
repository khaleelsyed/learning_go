package context

import "testing"

func TestBasicContextTimeOut(t *testing.T) {

	testCases := []struct {
		name      string
		apiCaller APICaller
		timeout   int
		want      APIResponse
	}{
		{
			name:      "slow response - 150ms timeout",
			apiCaller: NewSlowAPICaller(),
			timeout:   150,
			want:      youStupid,
		},
		{
			name:      "slow response - 400ms timeout",
			apiCaller: NewSlowAPICaller(),
			timeout:   400,
			want:      answerToEverything,
		},
		{
			name:      "fast response - 150ms timout",
			apiCaller: NewFastAPICaller(),
			timeout:   150,
			want:      answerToEverything,
		},
		{
			name:      "fast response - 50ms timout",
			apiCaller: NewFastAPICaller(),
			timeout:   50,
			want:      youStupid,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, _ := handleAPICall(test.timeout, test.apiCaller)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
