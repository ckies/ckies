package ckies

import (
	"bytes"
	"errors"
	"path"
	"text/template"
)

type Builder struct {
	Name    string
	Info    string
	Website string

	Options struct {
		ForceTables bool
	}

	Links struct {
		Policy   string
		Settings string
	}

	Cookies struct {
		Necessary   []Cookie
		Functional  []Cookie
		Performance []Cookie
		Marketing   []Cookie
	}
}

// Parse template to Markdown file
func (b *Builder) Parse(file string) (*Document, error) {
	data := Document{&bytes.Buffer{}}
	name := path.Base(file)

	tmpl, err := template.New(name).ParseFiles(file)

	if err != nil {
		return nil, errors.New("Unable to parse template: " + file)
	}

	if tmpl.Execute(data.Data, b) != nil {
		return nil, errors.New("Unable to render template: " + file)
	}

	return &data, nil
}

// NewBuilder returns a Builder with config
func NewBuilder(config Config, services map[string]Service) Builder {
	builder := Builder{
		Name:    config.Name,
		Info:    config.Info,
		Website: config.Website,
		Links:   config.Links,
		Options: config.Options,
	}

	merger := Merger{config, services}
	builder.Cookies.Necessary = merger.filterByType(CookieTypeNecessary)
	builder.Cookies.Functional = merger.filterByType(CookieTypeFunctional)
	builder.Cookies.Performance = merger.filterByType(CookieTypePerformance)
	builder.Cookies.Marketing = merger.filterByType(CookieTypeMarketing)

	return builder
}
