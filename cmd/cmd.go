package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/ckies/ckies/ckies"
	"github.com/spf13/cobra"
)

var Version string
var Config *ckies.Config
var Services map[string]ckies.Service

var cfgFile string
var cfgOutput string
var cfgServices string
var cfgTemplates string

var rootCmd = &cobra.Command{
	Use:   "ckies",
	Short: "A brief description of your application",
}

func abort(format string, a ...interface{}) {
	write("[Error] "+format+"\n", a...)
	os.Exit(1)
}

func write(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/ckies.toml)")
	rootCmd.PersistentFlags().StringVar(&cfgOutput, "output", "", "folder to store genrated files")
	rootCmd.PersistentFlags().StringVar(&cfgServices, "services", "", "prefix for service lookup")
	rootCmd.PersistentFlags().StringVar(&cfgTemplates, "templates", "", "prefix for template lookup")
}

func requireValidConfiguration() {
	// Config file is always required
	if Config == nil {
		abort("Unable to open config file %s", cfgFile)
	} else {
		write("Using config file %s\n", cfgFile)
	}

	// Output flag is always required
	if cfgOutput == "" {
		abort("Please configure destination using --output")
	} else {
		if _, err := os.Stat(cfgOutput); os.IsNotExist(err) {
			abort("Configured destination does not exist: %s", cfgOutput)
		}

		if handle, err := os.Stat(cfgOutput); err == nil && !handle.IsDir() {
			abort("Configured destination is not a directory: %s", cfgOutput)
		}
	}

	// Templates flag is always required
	if cfgTemplates == "" {
		abort("Please configure prefix for template lookup using --templates")
	} else {
		if _, err := os.Stat(cfgTemplates); os.IsNotExist(err) {
			abort("Configured template prefix does not exist: %s", cfgTemplates)
		}

		if handle, err := os.Stat(cfgTemplates); err == nil && !handle.IsDir() {
			abort("Configured template prefix is not a directory: %s", cfgTemplates)
		}
	}

	// Service prefix is only needed if config uses services
	if len(Config.Services) > 0 {
		if cfgServices == "" {
			abort("Please configure prefix for service lookup using --services")
		} else {
			if _, err := os.Stat(cfgServices); os.IsNotExist(err) {
				abort("Configured service prefix does not exist: %s", cfgServices)
			}

			if handle, err := os.Stat(cfgServices); err == nil && !handle.IsDir() {
				abort("Configured service prefix is not a directory: %s", cfgServices)
			}
		}

		for _, serviceKey := range Config.Services {
			service, err := ckies.GetServiceFromPrefix(serviceKey, cfgServices)

			if err != nil {
				abort("Unable to find service `%s` in service prefix: %s", service, cfgServices)
			}

			// Store service information in memory
			Services[serviceKey] = *service
		}
	}
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = path.Join("ckies.toml")
	}

	Config, _ = ckies.Load(cfgFile)
	Services = map[string]ckies.Service{}
}

// Execute handles the CLI command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
