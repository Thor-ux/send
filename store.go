package main

var suppressed = make(map[string]bool)
var sentToday = make(map[string]bool)

func isBounced(email string) bool {
        return suppressed[email]
}

func alreadySentToday(email string) bool {
        return sentToday[email]
}

func markSent(email string) {
        sentToday[email] = true
}

func addToSuppression(email string) {
        suppressed[email] = true
}