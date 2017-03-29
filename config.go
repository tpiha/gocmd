package main

type Config struct {
	Dev bool
}

func initConfig() *Config {
	conf := &Config{}
	return conf
}
