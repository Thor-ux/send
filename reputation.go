package main

import (
        "encoding/json"
        "os"
        "time"
)

type Reputation struct {
        Score      int    `json:"score"`
        Bounces    int    `json:"bounces"`
        Complaints int    `json:"complaints"`
        LastUpdate string `json:"last_update"`
}

const ReputationFile = "reputation.json"

func loadReputation() Reputation {
        f, err := os.Open(ReputationFile)
        if err != nil {
                return Reputation{Score: 100, LastUpdate: time.Now().UTC().Format("2006-01-02")}
        }
        defer f.Close()
        var r Reputation
        json.NewDecoder(f).Decode(&r)
        return r
}

func saveReputation(r Reputation) {
        f, _ := os.Create(ReputationFile)
        defer f.Close()
        json.NewEncoder(f).Encode(r)
}

func updateReputation(event string) {
        r := loadReputation()
        switch event {
        case "bounce":
                r.Bounces++
                r.Score -= 5
        case "complaint":
                r.Complaints++
                r.Score -= 15
        }
        if r.Score < 0 {
                r.Score = 0
		        }
        r.LastUpdate = time.Now().UTC().Format("2006-01-02")
        saveReputation(r)
}

func recoverReputation() {
        r := loadReputation()
        today := time.Now().UTC().Format("2006-01-02")
        if r.LastUpdate == today {
                return
        }
        r.Score += 5
        if r.Score > 100 {
                r.Score = 100
        }
        r.LastUpdate = today
        saveReputation(r)
}

func reputationDelay(base int) int {
        recoverReputation()
        r := loadReputation()
        switch {
        case r.Score >= 80:
                return base
        case r.Score >= 60:
                return base + 15
        case r.Score >= 40:
                return base + 30
        default:
                return base + 60
        }
}
