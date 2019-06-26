package main

import (
	"log"
	"net/http"
	"os"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "8080"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	writer.Header().Set("Access-Control-Allow-Origin", "*")

	// allow pre-flight headers
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	request.Write(writer)
}

// // RandomHandler responses with random int (prevent caching for tests)
// func RandomHandler(writer http.ResponseWriter, request *http.Request) {
// 	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")
// 	i := rand.Float64()
// 	request.Write(i)
// }

func main() {

	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe("0.0.0.0:"+getServerPort(), nil)
}
