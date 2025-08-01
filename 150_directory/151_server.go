// Moving this to 150 folder where the rest of stuff is.

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)
func main(){

	http.HandleFunc("/orders",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key:= "key.pem"

	// configure TLS

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	// enable http2

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}

	//HTTP 1.1 without tls
	// fmt.Println("Server is running on port:", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server", err)
	// }
}
// dependencies are all stored outside 
// the code base within the go path location.
//  openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365   
// Read more on TLS

// 152: HTTPS certificates - ssl/tls
// key.pem contains private key.
// cert.pem contains public key
// PEM: Privacy enhanced mail: cryptograpic stuff
// DER: Distinguished encoding rules
// PEM is 64 encoded. ASCII.  Footers and headers identify info they hold.
// Private key is in begin private key, end private key
// certificate is between begin cert nad end cert.
// .crt and .key files can be used instead of .pem files.

// 153: All about postman

// 154: Curl: 