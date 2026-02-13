package main

import (
        "encoding/json"
        "log"
        "net/http"
        "os"
        "time"
)

// =====================
// TYPES
// =====================
type BounceEvent struct {
        Email  string `json:"email"`
        Type   string `json:"type"`   // bounce | complaint
        Reason string `json:"reason"`
}

// =====================
// HANDLER
// =====================
func webhookHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
                http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
                return
        }
        defer r.Body.Close()

        var event BounceEvent
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
                http.Error(w, "invalid payload", http.StatusBadRequest)
                return
        }

        if event.Email == "" || event.Type == "" {
                http.Error(w, "missing fields", http.StatusBadRequest)
                return
        }

        now := time.Now().UTC().Format(time.RFC3339)

        switch event.Type {
        case "bounce":
                appendLine("bounce.log", now+" "+event.Email+" "+event.Reason+"\n")
                updateReputation("bounce")
				                log.Printf("📛 WEBHOOK BOUNCE: %s\n", event.Email)

        case "complaint":
                appendLine("complaint.log", now+" "+event.Email+" "+event.Reason+"\n")
                updateReputation("complaint")
                log.Printf("⚠️ WEBHOOK COMPLAINT: %s\n", event.Email)

        default:
                http.Error(w, "unknown event type", http.StatusBadRequest)
                return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
}

// =====================
// SERVER
// =====================
func startWebhook() {
        mux := http.NewServeMux()
        mux.HandleFunc("/webhook", webhookHandler)
        mux.HandleFunc("/unsubscribe", unsubscribeHandler)

        log.Println("📡 Webhook listening on :80/webhook")

        // Never use log.Fatal here — it kills the container
        if err := http.ListenAndServe(":80", mux); err != nil {
                log.Println("webhook server stopped:", err)
        }
}

// =====================
// FILE APPEND
// =====================
func appendLine(file, line string) {
        f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
                log.Println("file write error:", err)
				                return
        }
        defer f.Close()

        _, _ = f.WriteString(line)
}