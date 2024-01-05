package main

import (
	"github.com/asylvis5951/pokedexcli/internal/pokeapi"

	pokecache "github.com/asylvis5951/pokedexcli/internal/pokecache"
	"github.com/asylvis5951/pokedexcli/internal/pokedex"

	"fmt"
)

func commandExplore(cfg *config, cache *pokecache.Cache, pd *pokedex.Pdex, s string) error {
	// get location by id
	fmt.Printf("Exploring %v...\n", s)
	_, err := pokeapi.GetLocationInfo(s, cache)
	if err != nil {
		return err
	}
	return nil
}
