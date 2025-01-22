package main

import (
	mflag "github.com/namsral/flag"
)

type Config struct {
	Port  int
	Debug bool
}

func parseConf() *Config {
	cfg := &Config{}

	mflag.String(mflag.DefaultConfigFlagname, "", "path to config file")

	mflag.IntVar(&cfg.Port, "server-port", 8080, "port use for the server")

	mflag.BoolVar(&cfg.Debug, "debug", true, "use debug")

	mflag.Parse()

	return cfg
}
