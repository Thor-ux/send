package main

import (
        "encoding/json"
        "os"
        "sync"
        "time"
)

type SuppressionEntry struct {
        Reason    string
        Timestamp time.Time
}

var (
        suppressionFile = "suppression.json"
        suppressions    = make(map[string]SuppressionEntry)
        mu              sync.Mutex
)

func loadSuppressions() {
        data, err := os.ReadFile(suppressionFile)
        if err != nil {
                return
        }
        json.Unmarshal(data, &suppressions)
}

func saveSuppressions() {
        data, _ := json.MarshalIndent(suppressions, "", "  ")
        os.WriteFile(suppressionFile, data, 0644)
}

func suppress(email, reason string) {
        mu.Lock()
        defer mu.Unlock()

        suppressions[email] = SuppressionEntry{
                Reason:    reason,
                Timestamp: time.Now().UTC(),
        }

        saveSuppressions()
}

func isSuppressed(email string) bool {
	    mu.Lock()
        defer mu.Unlock()

        _, exists := suppressions[email]
        return exists
}