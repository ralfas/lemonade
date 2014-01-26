package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCreateListener(t *testing.T) {

	tests := []struct {
		Port            string
		ExpectedAddress string
	}{
		{"8080", "0.0.0.0:8080"},
		{"8082", "0.0.0.0:8082"},
	}

	for testIndex, test := range tests {

		listener := createListener(test.Port)

		if listener.Addr().String() != test.ExpectedAddress {
			t.Errorf("%d. Network address returned => %q, expected => %q", testIndex, listener.Addr().String(), test.ExpectedAddress)
		}
	}
}

func TestRequestHandler(t *testing.T) {

	tests := []struct {
		Request            http.Request
		ExpectedStatusCode int
	}{
		{http.Request{URL: &url.URL{Host: "0.0.0.0:8080", Path: "/"}}, 200},
		{http.Request{URL: &url.URL{Host: "0.0.0.0:8081", Path: "/"}}, 200},
		{http.Request{URL: &url.URL{Host: "0.0.0.0:8080", Path: "/sdfsdf"}}, 404},
	}

	for testIndex, test := range tests {

		recorder := httptest.NewRecorder()

		requestHandler(recorder, &test.Request)

		if recorder.Code != test.ExpectedStatusCode {
			t.Errorf("%d. Status Code returned => %d, expected => %d", testIndex, recorder.Code, test.ExpectedStatusCode)
		}
	}
}

func TestGenerateTwilML(t *testing.T) {

	tests := []struct {
		Message        string
		ExpectedTwilML string
	}{
		{"", `<?xml version="1.0" encoding="UTF-8"?><Response><Say voice="woman"></Say></Response>`},
		{"call me", `<?xml version="1.0" encoding="UTF-8"?><Response><Say voice="woman">call me</Say></Response>`},
		{"Pick up! There's an emergency.", `<?xml version="1.0" encoding="UTF-8"?><Response><Say voice="woman">Pick up! There's an emergency.</Say></Response>`},
	}

	for testIndex, test := range tests {

		twilML := generateTwilML(test.Message)

		if twilML != test.ExpectedTwilML {
			t.Errorf("%d. TwilML returned => %q, expected => %q", testIndex, twilML, test.ExpectedTwilML)
		}
	}
}

// TODO
func TestMakeCall(t *testing.T) {

	tests := []struct {
		APIPath   string
		From      string
		To        string
		TwilMLURL url.URL
	}{
		{"", "", "", url.URL{}},
	}

	for testIndex, test := range tests {

		testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path != test.APIPath {
				t.Errorf("%d. Request path => %q, expected => %q", testIndex, request.URL.Path, test.APIPath)
			}
		}))
		defer testServer.Close()

		apiURL := url.URL{Host: testServer.URL, Path: test.APIPath}

		makeCall(apiURL, test.From, test.To, test.TwilMLURL)
	}
}
