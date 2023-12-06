package config

import "github.com/kelseyhightower/envconfig"

type Environment struct {
	Port             int    `required:"true"`
	PostgresAddress  string `required:"true" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
	PostgresDatabase string `required:"true" split_words:"true"`
}

// Get returns the current environment.
func Get() (*Environment, error) {
	var e Environment
	err := envconfig.Process("", &e)
	return &e, err
}
