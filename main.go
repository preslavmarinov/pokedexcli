package main

import (
	"time"

	"github.com/preslavmarinov/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewCLient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
