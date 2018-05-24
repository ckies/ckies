package ckies

import (
	"bytes"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	blackfridayOptions = blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
)

// Document stores the built markdown document in memory
type Document struct {
	Data *bytes.Buffer
}

func (d *Document) Format(f Formatter) []byte {
	return f(d)
}

// Markdown return the document as Markdown
func (d *Document) Markdown() []byte {
	return d.Format(ToMarkdown)
}

// HTML returns the document as HTML
func (d *Document) HTML() []byte {
	return d.Format(ToHTML)
}
