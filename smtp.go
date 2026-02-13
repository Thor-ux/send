package main

import (
        "net/smtp"
)

func sendEmail(to, subject, body string) error {
        if isSuppressed(to) {
                return nil
        }

        msg := buildMessage(to, subject, body)

        auth := smtp.PlainAuth(
                "",
                cfg.SMTPUser,
                cfg.SMTPPass,
                cfg.SMTPHost,
        )

        err := smtp.SendMail(
                cfg.SMTPHost+":"+cfg.SMTPPort,
                auth,
                cfg.ReturnPath,
                []string{to},
                []byte(msg),
        )

        return err
}