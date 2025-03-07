package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var port *string

func init() {
	port = flag.String("port", "80", "port to serve on")
}

func main() {
	flag.Parse()
	handler := http.FileServer(http.Dir("./static"))

	addr := net.JoinHostPort("0.0.0.0", *port)
	log.Printf("Serving on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, handler))
}
