                                                package main

import (
        "fmt"
        "strings"
        "time"

        "github.com/google/uuid"
)

func sanitizeHeader(value string) string {
        value = strings.ReplaceAll(value, "\r", "")
        value = strings.ReplaceAll(value, "\n", "")
        return value
}

func buildMessage(to, subject, htmlBody string) string {

        // Prevent header injection
        subject = sanitizeHeader(subject)

        messageID := fmt.Sprintf("<%s@lhcontractorgroup.design>", uuid.NewString())

        unsubURL := fmt.Sprintf("%s/unsubscribe?email=%s", cfg.BaseURL, to)

        boundary := "mixed-" + uuid.NewString()

        plainText := stripHTML(htmlBody)

        return fmt.Sprintf(
                "From: %s\r\n"+
                        "To: %s\r\n"+
                        "Subject: %s\r\n"+
                        "Date: %s\r\n"+
                        "Message-ID: %s\r\n"+
                        "Return-Path: %s\r\n"+
                        "Precedence: bulk\r\n"+
                        "List-Unsubscribe: <%s>\r\n"+
                        "List-Unsubscribe-Post: List-Unsubscribe=One-Click\r\n"+
                        "X-Mailer: LHMailer/1.0\r\n"+
                        "MIME-Version: 1.0\r\n"+
                        "Content-Type: multipart/alternative; boundary=%s\r\n"+
                        "\r\n"+
                        "--%s\r\n"+
                        "Content-Type: text/plain; charset=UTF-8\r\n"+
                        "Content-Transfer-Encoding: 8bit\r\n"+
						"\r\n"+
                        "%s\r\n"+
                        "\r\n"+
                        "--%s--\r\n",
                cfg.FromAddress,
                to,
                subject,
                time.Now().UTC().Format(time.RFC1123Z),
                messageID,
                cfg.ReturnPath, // MUST match SES verified domain
                unsubURL,
                boundary,
                boundary,
                plainText,
                boundary,
                htmlBody,
                boundary,
        )
}