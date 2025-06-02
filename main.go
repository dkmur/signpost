package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
    "time"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml/v2"
)

var pokemonTemplate *template.Template
var pokestopTemplate *template.Template
var gymTemplate *template.Template

func main() {
	tomlFile, err := os.Open("config.toml")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	// defer the closing of our tomlFile so that we can parse it later on
	defer tomlFile.Close()

	byteValue, _ := io.ReadAll(tomlFile)

	err = toml.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
	config.TimestampFormat = "2006-01-02 15:04:05"

	//	templateStr := "https://maps.google.com/maps?q={{.lat}},{{.lon}}"

	pokemonTemplate = template.New("pokemon")
	pokestopTemplate = template.New("pokestop")
	gymTemplate = template.New("gym")

	for _, t := range config.Pokemon {
		pokemonTemplate, err = pokemonTemplate.New(t.Name).Parse(t.Url)
		if err != nil {
			panic(err)
		}
	}

	for _, t := range config.Pokestop {
		pokestopTemplate, err = pokestopTemplate.New(t.Name).Parse(t.Url)
		if err != nil {
			panic(err)
		}
	}

	for _, t := range config.Gym {
		gymTemplate, err = gymTemplate.New(t.Name).Parse(t.Url)
		if err != nil {
			panic(err)
		}
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/pokemon/:pokemon_id/:template", GetPokemon)
	r.GET("/pokestop/:pokestop_id/:template", GetPokestop)
	r.GET("/gym/:gym_id/:template", GetGym)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,
	}
	fmt.Printf("%s [] Starting server on port %d\n", time.Now().Format(config.TimestampFormat), config.Port)
	srv.ListenAndServe()
}
