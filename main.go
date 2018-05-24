package main

import "github.com/ckies/ckies/cmd"

var version string

func main() {
	if version != "" {
		cmd.Version = version
	} else {
		cmd.Version = "0.0.0-dev"
	}

	cmd.Execute()
}
