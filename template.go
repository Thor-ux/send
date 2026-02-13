package main

import (
        "bytes"
        "html/template"
)

func renderTemplate(path string, data any) (string, error) {

        tmpl, err := template.ParseFiles(path)
        if err != nil {
                return "", err
        }

        var buf bytes.Buffer

        err = tmpl.Execute(&buf, data)
        if err != nil {
                return "", err
        }

        return buf.String(), nil
}