package tinyml

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	autocorrect "github.com/huacnlee/go-auto-correct"
)

// ToHTML parse TinyML format to HTML
func ToHTML(raw string) (out string, err error) {
	w := &bytes.Buffer{}
	l := NewLexer(bytes.NewBufferString(raw))
	out = raw
	tagName := ""

	for {
		token, data := l.Next()

		switch token {
		case ErrorToken:
			if l.Err() == io.EOF {
				out = w.String()
				paragraphs := strings.Split(out, "\n\n")

				w := &bytes.Buffer{}
				for _, p := range paragraphs {
					// p = strings.ReplaceAll(p, "\n", "<br />")
					w.WriteString("<p>" + strings.TrimSpace(p) + "</p>")
				}

				out, err = autocorrect.FormatHTML(w.String())
				if err != nil {
					return w.String(), err
				}
				return out, nil
			}
		case BreakLineToken:
			w.Write([]byte("<br />"))
			continue
		case StartTagToken:
			tag := string(data)
			if tag == "[st]" {
				tagName = "security"
			}
			continue
		case EndTagToken:
			tagName = ""
			continue
		}

		if tagName != "" {
			if tagName == "security" {
				value := string(data)
				counterID, name, err := parseSecurityTag(string(data))
				if err != nil {
					if _, err := w.Write(data); err != nil {
						return out, err
					}

					continue
				}

				name = fmt.Sprintf(" <span class=\"security-tag\" value=\"%s\" data-id=\"%s\">%s</span> ", value, counterID, name)
				if _, err := w.WriteString(name); err != nil {
					return out, err
				}
			}

			continue
		}

		val := bytes.Trim(data, " ")
		if _, err := w.Write(val); err != nil {
			return out, err
		}
	}
}
