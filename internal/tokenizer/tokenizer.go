package tokenizer

import (
	"golang.org/x/net/html"
	"strings"
)

func ParseHTML(s string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(s))
	var textContent string
	inBody := false
	inScript := false

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			// End of the document
			return textContent
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "body" {
				inBody = true
			} else if token.Data == "script" {
				inScript = true
			}
		case html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "body" {
				inBody = false
			} else if token.Data == "script" {
				inScript = false
			}
		case html.TextToken:
			if inBody && !inScript {
				// Append text content
				text := strings.TrimSpace(string(tokenizer.Text()))
				if len(text) > 0 {
					textContent += text + " "
				}
			}
		}
	}
}