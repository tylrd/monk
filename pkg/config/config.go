package config

type Config struct {
	Global Global `toml:"global"`
}

type Global struct {
	Verbose bool `toml:"verbose"`
}

func Default() *Config {
	return &Config{
		Global: Global{
			Verbose: true,
		},
	}
}
