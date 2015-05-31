package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type CircuitConnection struct {
	Ws *websocket.Conn
}

type PendingRequest struct {
	Request *http.Request
	// the CircuitConnection that is currently handling this reuest, or nil
	HandlingConnection *CircuitConnection
}

type Circuit struct {
	CircuitConnections []CircuitConnection
	PendingRequests    []PendingRequest
	OperatorKey        string
	ExitKey            string
}

// Handles a new circuit websocket connection, through which requests may be proxied.
func circuitHandler(ws *websocket.Conn) {
	fmt.Println("Got WS connection ", ws)
}

// Handles a request from the browser, to be proxied if possible.
func browserHandler(w http.ResponseWriter, r *http.Request) {
	// placeholder:
	http.Redirect(w, r, "/__xssrc__/browser", 302)
}

func main() {
	http.Handle("/__xssrc__/circut/connection", websocket.Handler(circuitHandler))
	http.Handle("/__xssrc__/",
		http.StripPrefix("/__xssrc__", http.FileServer(http.Dir("./client"))))
	http.Handle("/", http.HandlerFunc(browserHandler))

	fmt.Println("Starting HTTP server")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic("Server error: " + err.Error())
	}
}

/*
TO DO
-
- After getting a request, wait some amount of time for a Circuit to
   become available to handle it, then send a 504 Gateway Timeout.
 - Always send 504 Gateway Timeout with a Refresh header
 - Support GET and POST requests, and arbitrary responses
 - Use an XSSRC- header to indicate requests that are not proxied.
 - Add very visible console logging on circuit clients, including the IPs of connecting
   browsers. Maybe persist some information (in a compact form) in localStorage, to enable a sort-of auditing.
 - Alway use the older circuit that is stil responsive, to avoid uses from switching between sessions (though
   that won't happen in ideal circuimstances anyway).
*/
