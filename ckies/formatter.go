package ckies

import (
	"github.com/shurcooL/markdownfmt/markdown"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Formatter func(d *Document) []byte

func ToMarkdown(d *Document) []byte {
	output, _ := markdown.Process("", d.Data.Bytes(), nil)

	return output
}

func ToHTML(d *Document) []byte {
	return blackfriday.Run(
		d.Markdown(),
		blackfriday.WithExtensions(blackfridayOptions),
	)
}
