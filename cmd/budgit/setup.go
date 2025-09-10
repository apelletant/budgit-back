package main

import "github.com/apelletant/logger"

type dependencies struct {
	log *logger.Logger
}

func setupDeps(cfg *Config) (*dependencies, error) {
	logger := logger.NewLogger(cfg.Debug)

	return &dependencies{
		log: logger,
	}, nil
}
