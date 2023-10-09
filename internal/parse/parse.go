package parse

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
	"webscraper/internal/errors"
)

type Root struct {
	Pointer   *html.Node
	NodeValue string
	Error     error
}

func (r Root) FullText() string {
	var buf bytes.Buffer

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.TextNode {
			buf.WriteString(strings.TrimSpace(n.Data))
			buf.WriteString(" ")
		}
		if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
			f(n.FirstChild)
		}
		if n.NextSibling != nil {
			f(n.NextSibling)
		}
	}

	f(r.Pointer.FirstChild)

	return buf.String()
}

func HTMLParse(s string) Root {
	r, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return Root{Error: errors.NewError(errors.ErrUnableToParse, "unable to parse the HTML")}
	}

	for r.Type != html.ElementNode {
		switch r.Type {
		case html.DocumentNode:
			r = r.FirstChild
		case html.DoctypeNode:
			r = r.NextSibling
		case html.CommentNode:
			r = r.NextSibling
		}
	}
	return Root{Pointer: r, NodeValue: r.Data}
}
