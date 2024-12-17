// client.go
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
    tlsConfig := &tls.Config{
        InsecureSkipVerify: true,
    }

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: tlsConfig,
        },
    }

    // Request the index.html file
    resp, err := client.Get("https://localhost:8443/index.html")
    if err != nil {
        log.Fatalf("failed to get: %s", err)
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}

	fmt.Println(string(body))
}
