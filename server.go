// server.go
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Load server certificate and key
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Fatalf("failed to load key pair: %s", err)
    }

    // Create a TLS config
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
    }

    // Create a file server handler
    fileServer := http.FileServer(http.Dir("./files"))

    // Create a new HTTP server with TLS
    server := &http.Server{
        Addr:      ":8443",
        Handler:   fileServer,
        TLSConfig: tlsConfig,
    }

	fmt.Println("Starting server on https://localhost:8443")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
