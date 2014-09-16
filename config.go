package main

type Config struct {
	DataPath      string
	SCM           string
	SCMAutoCommit bool
	SCMAutoPush   bool
}

func NewConfig() Config {
	config := new(Config)
	return *config
}
