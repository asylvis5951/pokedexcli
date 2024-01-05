package main

import (
	"fmt"
	"github.com/asylvis5951/pokedexcli/internal/pokeapi"
	"github.com/asylvis5951/pokedexcli/internal/pokecache"
	"github.com/asylvis5951/pokedexcli/internal/pokedex"

)

func commandCatch(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, s string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", s)
	// Get pokemon info
	err := pokeapi.GetPokemon(s, cache, p)
	if err != nil {
		return err
	}
	return nil
}
