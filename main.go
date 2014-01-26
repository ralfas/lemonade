package main

import (
	//"github.com/mitchellh/goamz/s3"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// mux := http.NewServeMux()
// mux.HandleFunc("/", requestHandler})

func createListener(port string) (listener net.Listener) {

	listener, err := net.Listen("tcp", strings.Join([]string{":", port}, ""))
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
}

func generateTwilML(message string) (twilML string) {

	const prefix string = `<?xml version="1.0" encoding="UTF-8"?><Response><Say voice="woman">`
	const suffix string = `</Say></Response>`

	return strings.Join([]string{prefix, message, suffix}, "")
}

func makeCall(apiURL url.URL, fromPhone string, toPhone string, twilMLURL url.URL) {

}

func main() {

	// configuration

	// Twilio credentials
	// AWS credentials
	// S3 bucket

	// create + run server to handle notification events
	// receives a phone number and the message

	// creates an object in S3 with a temporary random name containing the message in a TwilML
	// object is set to expire in 30 minutes

	// generates a Signed URL for S3 object and passes to Twilio
}
