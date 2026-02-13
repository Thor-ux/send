package main

import "os"

type Config struct {
        SMTPHost    string
        SMTPPort    string
        SMTPUser    string
        SMTPPass    string
        FromAddress string
        ReturnPath  string
        BaseURL     string
}

func LoadConfig() Config {
        return Config{
                SMTPHost:    os.Getenv("SMTP_HOST"),
                SMTPPort:    os.Getenv("SMTP_PORT"),
                SMTPUser:    os.Getenv("SMTP_USER"),
                SMTPPass:    os.Getenv("SMTP_PASS"),
                FromAddress: "LH Contractor Group <hello@lhcontractorgroup.design>",
                ReturnPath:  "bounce@lhcontractorgroup.design",
                BaseURL:     "https://lhcontractorgroup.design",
        }
}