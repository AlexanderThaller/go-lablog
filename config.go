package main

type Config struct {
	DataPath string
}

func NewConfig() Config {
	config := new(Config)
	return *config
}
