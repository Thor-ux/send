package main

import (
        "bufio"
        "os"
)

func loadRecipients() []string {
        file, err := os.Open("recipients.txt")
        if err != nil {
                return nil
        }
        defer file.Close()

        var recipients []string
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
                recipients = append(recipients, scanner.Text())
        }

        return recipients
}