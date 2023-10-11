package tokenizer

import (
	"golang.org/x/net/html"
	"strings"
)

func ParseHTML(s string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(s))
	var textContent string
	inScript := false
	inStyle := false

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			// End of the document
			return textContent
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "script" {
				inScript = true
			} else if token.Data == "style" {
				inStyle = true
			}
		case html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "script" {
				inScript = false
			} else if token.Data == "style" {
				inStyle = false
			}
		case html.TextToken:
			if !inScript && !inStyle {
				// Append text content
				text := strings.TrimSpace(strings.ToLower(string(tokenizer.Text())))
				if len(text) > 0 {
					textContent += text + " "
				}
			}
		}
	}
}
