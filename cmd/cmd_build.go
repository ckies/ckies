package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/ckies/ckies/ckies"
	"github.com/spf13/cobra"
)

var (
	cfgFormat string
	files     = []string{
		"policy.md",
		"settings.md",
	}
	exports = map[string]struct {
		formatter ckies.Formatter
		extension string
	}{
		"html":     {ckies.ToHTML, "html"},
		"markdown": {ckies.ToMarkdown, "md"},
	}
)

func init() {
	buildCmd.Flags().StringVar(&cfgFormat, "format", "html", "Export format")

	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build files",
	Run: func(cmd *cobra.Command, args []string) {
		requireValidConfiguration()

		builder := ckies.NewBuilder(*Config, Services)

		if _, ok := exports[cfgFormat]; !ok {
			abort("Unsupported export format: %s", cfgFormat)
		}

		for _, file := range files {
			doc, err := builder.Parse(path.Join(cfgTemplates, file))
			format := exports[cfgFormat]

			if err != nil {
				abort("Failed to render template %q: %s", file, err)
			}

			name := strings.TrimSuffix(file, path.Ext(file)) + "." + format.extension
			data := doc.Format(format.formatter)

			ioutil.WriteFile(path.Join(cfgOutput, name), data, 0644)

			fmt.Printf("Wrote %s\n", path.Join(cfgOutput, name))
		}
	},
}
