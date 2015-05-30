///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

type CircuitConnection struct {
    Ws *ws.Conn
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

func circuitHandler(ws *webSocket.Conn) {
	fmt.Println("Got connection ", ws)
}

func main() {
	http.Handle("/__xssrc__/circut-connection", websocket.Handler(circuitHandler))
	http.Handle("/__xssrc__", http.FileServer(http.Dir("static"))

	fmt.Println(&Circuit{})

	fmt.Println("WebSockets Version", websocket.SupportedProtocolVersion)

	// Intentionally loopback-only. Do not change this to 0.0.0.0.
	err := http.ListenAndServe("127.0.0.1:8064", nil)
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
   controllers. Maybe persist some information (in a compact form) in localStorage, to enable a sort-of auditing.
 - Alway use the older circuit that is stil responsive, to avoid uses from switching between sessions (though
   that won't happen in ideal circuimstances anyway).
*/
