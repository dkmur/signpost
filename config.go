package main

type Config struct {
	Port            int                  `toml:"port"`
	Golbat          golbatConfiguration  `toml:"golbat"`
	Pokemon         []templateDefinition `toml:"pokemon"`
	Pokestop        []templateDefinition `toml:"pokestop"`
	Gym             []templateDefinition `toml:"gym"`
	TimestampFormat string
}

type golbatConfiguration struct {
	Url         string `toml:"url"`
	ApiPassword string `toml:"api_password"`
}

type templateDefinition struct {
	Url  string `toml:"url"`
	Name string `toml:"name"`
}

var config Config
