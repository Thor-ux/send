package main

import (
        "log"
        "net/http"
)

var cfg Config

func main() {
        log.Println("🚀 Mailer starting")

        cfg = LoadConfig()

        loadSuppressions()

        go startHTTPServer()
        go startWarmupLoop()

        select {} // block forever
}

func startHTTPServer() {
        mux := http.NewServeMux()
        mux.HandleFunc("/unsubscribe", unsubscribeHandler)

        log.Println("📡 HTTP server listening on :80")
        err := http.ListenAndServe(":80", mux)
        if err != nil {
                log.Fatal(err)
        }
}