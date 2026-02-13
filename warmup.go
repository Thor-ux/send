package main

import (
        "log"
        "time"
)

const (
        DailyLimit       = 300
        PauseBetweenSend = 5 * time.Second
)

func startWarmupLoop() {

        for {
                log.Println("🌅 Starting daily warmup cycle")

                sentToday := 0

                recipients := loadRecipients()

                for _, r := range recipients {

                        // Stop if daily limit reached
                        if sentToday >= DailyLimit {
                                log.Println("📛 Daily limit reached")
                                break
                        }

                        // Skip bounced emails
                        if isBounced(r) {
                                log.Printf("⚠️ Skipping bounced: %s", r)
                                continue
                        }

                        // Skip already sent (avoid duplicates)
                        if alreadySentToday(r) {
                                log.Printf("⚠️ Already sent today: %s", r)
                                continue
                        }

                        // Render HTML
                        body, err := renderTemplate("templates/email.html", map[string]any{
                                "Email": r,
                        })
                        if err != nil {
							                                log.Println("Template error:", err)
                                continue
                        }

                        // Send email
                        err = sendEmail(r, "ADD SUBJECT", body)
                        if err != nil {
                                log.Printf("❌ FAILED: %s (%v)", r, err)
                                continue
                        }

                        markSent(r)
                        sentToday++

                        log.Printf("✅ SENT: %s (%d/%d)", r, sentToday, DailyLimit)

                        time.Sleep(PauseBetweenSend)
                }

                log.Println("🌙 Sleeping until next day...")
                time.Sleep(24 * time.Hour)
        }
}