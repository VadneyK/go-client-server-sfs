# go-client-server-sfs


### Step 1: Generate TLS Certificates

For testing purposes, you can generate a self-signed certificate using OpenSSL:

**openssl** **req** **-x509** **-newkey** **rsa:4096** **-keyout** **server.key** **-out** **server.crt** **-days** **365** **-nodes**

### Step 2: Create a Directory for Files

Create a directory named files in the same directory as your Go files. Place any files you want to serve in this directory.

### Step 3: Run the Server and Client

* Run the server:

  **   **go** **run** server.go
* In a separate terminal, run the client:

  **   **go** **run** client.go

### Explanation

* TLS Configuration: The server and client are configured to use TLS for secure communication. In a production environment, you should verify the server's certificate on the client side.
* File Server: The server uses http.FileServer to serve files from the ./files directory.
* Client: The client makes a secure request to the server and prints the response.

This is a basic setup. For a production system, consider adding authentication, authorization, and more robust error handling.
