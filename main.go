// Package main is an entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goaltools/iniflag"
)

var (
	addr = flag.String("http.addr", ":8080", "HTTP address the app must listen on")

	// Parameters for using HTTPS mode.
	tls  = flag.Bool("tls.enable", false, "Start application using TLS")
	cert = flag.String("tls.cert", "", "Path to the TLS certificate file")
	key  = flag.String("tls.key", "", "Path to the TLS key file")
)

// main parses INI configuration files making their values
// available to the flags of the application, prepares a handler,
// and starts a new server.
func main() {
	// Parse INI configuration files to make their values
	// available as flags.
	err := iniflag.Parse("config/app.ini")
	assertNil(err)

	// Allocate and run a new server. Depending on the
	// configuration use either HTTP or HTTPS.
	s := &http.Server{
		Addr: *addr,
	}
	switch *tls {
	case true:
		err = s.ListenAndServeTLS(*cert, *key)
	default:
		err = s.ListenAndServe()
	}
	assertNil(err)
}

// assertNil gets an error as an argument and makes sure
// it is nil. If it isn't, it terminates the program.
func assertNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
