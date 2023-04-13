package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// ToHTML converts markdown content to HTML.
func ToHTML(content []byte) ([]byte, error) {
	md := goldmark.New(goldmark.WithExtensions(extension.GFM))

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
