package main

import (
	"fmt"
	pokecache "github.com/asylvis5951/pokedexcli/internal/pokecache"

	"github.com/asylvis5951/pokedexcli/internal/pokeapi"
	"github.com/asylvis5951/pokedexcli/internal/pokedex"
)

func commandMapf(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, s string) error {
	lr, err := pokeapi.GetLocations(cfg.nextUrl, cache)
	if err != nil {
		return err
	}
	cfg.nextUrl = &lr.Next
	cfg.prevUrl = &lr.Previous
	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, s string) error {
	if cfg.prevUrl == nil {
		fmt.Println("you're on the first page of locations")
		return nil
	}
	lr, err := pokeapi.GetLocations(cfg.prevUrl, cache)
	if err != nil {
		return err
	}
	cfg.nextUrl = &lr.Next
	cfg.prevUrl = &lr.Previous
	return nil
}
