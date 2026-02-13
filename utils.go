package main

import "strings"

func stripHTML(input string) string {
        replacer := strings.NewReplacer(
                "<br>", "\n",
                "<br/>", "\n",
                "<br />", "\n",
                "</p>", "\n",
                "<p>", "",
        )
        return replacer.Replace(input)
}