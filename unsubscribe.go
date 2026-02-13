package main

import (
        "log"
        "net/http"
)

func unsubscribeHandler(w http.ResponseWriter, r *http.Request) {

        email := r.URL.Query().Get("email")
        if email == "" {
                http.Error(w, "missing email", 400)
                return
        }

        suppress(email, "unsubscribe")

        log.Printf("📭 UNSUBSCRIBE: %s ip=%s ua=%s",
                email,
                r.RemoteAddr,
                r.UserAgent(),
        )

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("You have been unsubscribed."))
}