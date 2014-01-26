package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCreateListener(t *testing.T) {

	tests := []struct {
		Port string
	}{
		{"8080"},
		{"8082"},
	}

	for testIndex, test := range tests {

		listener := createListener(test.Port)

		_, listenerPort, _ := net.SplitHostPort(listener.Addr().String())
		if listenerPort != test.Port {
			t.Errorf("%d. Network address returned => %q, expected => %q", testIndex, listenerPort, test.Port)
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
// https://api.twilio.com
func TestMakeCall(t *testing.T) {

	tests := []struct {
		APIPath string
		From    string
		To      string
		Url     string
	}{
		{"/2010-04-01/Accounts/7788678767/Calls", "+0202020202", "+10101010101", "https://example.com/signedURLgoesHere"},
		{"/2010-04-01/Accounts/123123123/Calls", "+0303030495", "+45905059544", "http://www.example2.com/signedURLgoesHere"},
	}

	testsRun := 0

	for testIndex, test := range tests {

		testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path != test.APIPath {
				t.Errorf("%d. Request path => %q, expected => %q", testIndex, request.URL.Path, test.APIPath)
			}
			if request.FormValue("From") != test.From {
				t.Errorf("%d. 'From' value => %q, expected => %q", testIndex, request.FormValue("From"), test.From)
			}
			if request.FormValue("To") != test.To {
				t.Errorf("%d. 'To' value => %q, expected => %q", testIndex, request.FormValue("To"), test.To)
			}
			if request.FormValue("Url") != test.Url {
				t.Errorf("%d. 'Url' value => %q, expected => %q", testIndex, request.FormValue("Url"), test.Url)
			}
			testsRun += 1
		}))
		defer testServer.Close()

		APIUrl, _ := url.Parse(testServer.URL)
		APIUrl.Path = test.APIPath

		makeCall(APIUrl.String(), test.From, test.To, test.Url)
	}

	if testsRun != cap(tests) {
		t.Errorf("Tests run => %d, expected => %d", testsRun, cap(tests))
	}
}
