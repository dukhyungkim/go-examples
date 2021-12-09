package config

import "github.com/jessevdk/go-flags"

type DBOptions struct {
	ConfigPath string `long:"config" default:"config.yml" description:"path to config file"`

	MariaDB   bool `long:"mariadb" description:"use MariaDB"`
	Postgres  bool `long:"postgres" description:"use PostgreSQL"`
	Migration bool `long:"migration" description:"auto migration"`
}

func ParseDBFlags() (*DBOptions, error) {
	var options DBOptions
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
