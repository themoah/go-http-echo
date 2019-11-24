package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/likexian/whois-go"
)

// DefaultPort is the default port to use if once is not specified by the PORT environment variable
const DefaultPort = "8080"

func getServerPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// EchoHandler echos back the request as a response
func EchoHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Echoing back request made to " + r.URL.Path + " to client (" + r.RemoteAddr + ")")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// allow pre-flight headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	r.Write(w)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// randomHandler responses with random int (prevent caching for tests)
func randomHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Echoing back request made to " + r.URL.Path + " to client (" + r.RemoteAddr + ")")
	i := rand.Float64()
	iStr := fmt.Sprintf("%f", i)
	fmt.Fprintf(w, iStr)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar")
}

func ipWhoisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request route: " + r.URL.Path)
	t := r.URL.Path
	requestedIP := strings.Replace(t, "/ip/", "", -1)
	// i := net.ParseIP(requestedIP)
	// log.Println("Requsted IP: " + requestedIP)
	result, err := whois.Whois(requestedIP)
	if err == nil {
		fmt.Fprintf(w, result)
	}
}

func main() {

	log.Println("starting server, listening on port 0.0.0.0:" + getServerPort())

	http.HandleFunc("/", EchoHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/ip/", ipWhoisHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+getServerPort(), nil))
}
