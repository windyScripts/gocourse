// Moving this to 150 folder where the rest of stuff is.

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)
func main(){

	http.HandleFunc("/orders",func(w http.ResponseWriter, r *http.Request){
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request){
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key:= "key.pem"

	// configure TLS

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, // enforce mtls
		ClientCAs: loadClientCAs(),
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

func logRequestDetails( r *http.Request){
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version:", httpVersion)

	if r.TLS != nil{
		tlsVersion := getTLSVersionName((r.TLS.Version))
		fmt.Println("Received request with TLS version: ",tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version."
	}
}
// curl: curl -v -k https://localhost:3000/orders
// v for verbose, k to skip certificate related issues
// ALPN application layer protocol negotiation

/* 
155 http2/https/http, tls handshake

http2 doesn't necessarily require encryption
browsers have decided to support http2 with https, called http2 over TLS
h2 over tls

grpc requires http2
	can be used without https, but https is recommended for privacy and data integrity.

GRPC does not work with browsers currently.

http1.1: new connections per request. Persistent connections, require keep alive. If a persistent connection isn't 
for http1.1, each request has a tcp handshake overhead.

TLS handshake
persistent connections

HTTP2:
multiplexing
single tcp connection
prioritization of streams, allowing prioritization of requests
reduced latency

TLS handshake:

Client hello
Server hello (sends digital certificate, with public key to client)
Client validates certificate against certificate authorities
Two finish statements, one in and one out. In between them is a cypher exchange.

SSL: secure sockets layer
TLS: Transport layer security

SSL is deprecated now. last was in 1996
TLS first in 1999, 1.3 in 2018. 1.2 and 1.3 are in use

session keys created using pre-master secret are used to encrypt data in transport.

you can upload a cert to postman, allowing trust for cert like it was from a CA.

*/

/* 156
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf
^ for creating key and certificate

mutualTLS requires both client and server to authenticate each other
used in internal comms between microservices etc.


Review this(156) or 161:mTLS and Postman Settings
*/

func loadClientCAs() *x509.CertPool{
clientCAs := x509.NewCertPool()
caCert, err := os.ReadFile("cert.pem")
if err != nil{
	log.Fatalln("Could not load client CA:",err)
}
clientCAs.AppendCertsFromPEM(caCert)
return clientCAs
}

/* 
157:

http2 has TLS overhead
http2 has more features.

covers setting up http1.1:
	TLSNextProto: make(map[string]fund(*http.Server, *tls.conn, http.Handler))

brew install nghttp2
	This installs h2load for benchmarking.

http1 : 1 mil requests in 3 sec
http2: 1 mil in 6 seconds.

total volume of traffic is much less, however.

	*/

