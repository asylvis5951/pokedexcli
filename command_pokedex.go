package main

import (
	// "fmt"
	"github.com/asylvis5951/pokedexcli/internal/pokecache"
	"github.com/asylvis5951/pokedexcli/internal/pokedex"
	// "encoding/json"
)


func commandPdex(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, name string) error {
	p.GetAll()
	return nil
}
