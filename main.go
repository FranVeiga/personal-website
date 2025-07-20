package main

import (
	"fmt"
	"log"
	"net/http"
    "os"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Pong!"))
}

func MyErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
}

func main() {
    port, found := os.LookupEnv("PORT") 
    if !found {
        port = "8080"
    }
    fs := http.FileServer(http.Dir("./public"))

    http.Handle("/", fs)
    http.Handle("/ping", http.HandlerFunc(MyHandler))
    http.Handle("/error", http.HandlerFunc(MyErrorHandler))

    log.Println("Listening on port", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
        log.Fatal("Failed to start HTTP server:", err)
    }
}
