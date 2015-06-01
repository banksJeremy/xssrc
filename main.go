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

type Server struct {
	http.Server
	Circuit
	mux *http.ServeMux
}

func NewServer() (s *Server) {
	s = &Server{}

	s.mux = http.NewServeMux()
	s.mux.Handle("circuit.xssrc.com:8080/socket", websocket.Handler(s.serveCircuitSocket))
	s.mux.Handle("circuit.xssrc.com:8080/", http.FileServer(http.Dir("./client/circuit")))
	s.mux.Handle("localhost:8080/__xssrc__/", http.StripPrefix("/__xssrc__",
		http.FileServer(http.Dir("./client/browser"))))
	s.mux.HandleFunc("localhost:8080/", s.serveBrowserRequest)
	s.mux.HandleFunc("/", s.serveUnexpectedRequest)

	s.Handler = s
	s.Addr = "0.0.0.0:8080"

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request, delegating to mux", r.URL, r.Header)
	s.mux.ServeHTTP(w, r)
}

func (s *Server) serveBrowserRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got browser request", r)
}

func (s *Server) serveCircuitSocket(ws *websocket.Conn) {
	fmt.Println("Got WebSocket connection for circuit")
}

func (s *Server) serveUnexpectedRequest(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/__xssrc__/", 302)
}

func main() {
	server := NewServer()
	fmt.Println("Starting HTTP server")
	err := server.ListenAndServe()
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
