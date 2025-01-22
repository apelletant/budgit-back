package main

type dependencies struct{}

func setupDeps(_ *Config) (*dependencies, error) {
	return &dependencies{}, nil
}
