///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

type Circuit struct {
	circuitConnections []string // websocket connection
	operatorKey        string
	exitKey            string
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
