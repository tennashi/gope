package config

import "github.com/spf13/pflag"

// Flags represent values of the command line flags.
type Flags struct {
	ConfigFile  string
	ProjectFile string
	Port        string
}

// InitFlags initialize flags.
func InitFlags() *Flags {
	flags := &Flags{}
	pflag.StringVarP(&flags.ConfigFile, "config", "c", "./config.toml", "config file")
	pflag.StringVarP(&flags.ProjectFile, "project", "p", "./project.toml", "project file")
	pflag.StringVarP(&flags.Port, "port", "P", "1323", "port")
	pflag.Parse()
	return flags
}
