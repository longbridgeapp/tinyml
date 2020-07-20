package tinyml

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	autocorrect "github.com/huacnlee/go-auto-correct"
)

// ToText parse TinyML format to PlainText
func ToText(raw string) (out string, err error) {
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
				out = strings.TrimSpace(out)
				out = autocorrect.Format(out)
				return out, nil
			}
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
				_, name, err := parseSecurityTag(string(data))
				if err != nil {
					if _, err := w.Write(data); err != nil {
						return out, err
					}

					continue
				}

				name = fmt.Sprintf(" %s ", name)
				if _, err := w.WriteString(name); err != nil {
					return out, err
				}
			}

			continue
		}

		if _, err := w.Write(data); err != nil {
			return out, err
		}

	}

}
