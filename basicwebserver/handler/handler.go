package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	rootReqTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "basicwebserver_root_request_total",
		Help: "The total number of requests /",
	})
	rootHelloTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "basicwebserver_hello_request_total",
		Help: "The total number of requests /hello",
	})
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	fmt.Printf("got / request. body:\n%s\n", body)

	rootReqTotal.Inc()

	io.WriteString(w, "This is my website!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	rootHelloTotal.Inc()
	io.WriteString(w, "Hello, HTTP!\n")
}

func GetTest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello Page</title>
		</head>
		<body>
			<h1>Hello, HTTP!</h1>
			<p>This is a simple HTML page served by Go.</p>
		</body>
		</html>
	`)
}
