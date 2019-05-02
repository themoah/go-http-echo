package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
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

func randomHandler(writer http.ResponseWriter, request *http.Request) {
	// log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")
	i := rand.Float64()
	s := fmt.Sprintf("%f", i)
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, s)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("test test test" + r.RemoteAddr)
	w.Header().Set("Foo", "Bar")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello, World!")
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tHeapAlloc = %v ", m.HeapAlloc)
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func printSometimes() {
	for x := 0; x < 99; x++ {
		fmt.Println("hello" + fmt.Sprintf("%v", x))
		printMemUsage()
		time.Sleep(5 * time.Second)
	}
}

func main() {

	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", EchoHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/random", randomHandler)
	// go printSometimes()
	http.ListenAndServe(":"+getServerPort(), nil)
}
