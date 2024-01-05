package main

import (
	pokecache "github.com/asylvis5951/pokedexcli/internal/pokecache"
	pokedex "github.com/asylvis5951/pokedexcli/internal/pokedex"
	"time"
)

func main() {
	pokedex := pokedex.NewPokedex()
	cfg := &config{}
	cache := pokecache.NewCache(30 * time.Second)
	startRepl(cfg, cache, pokedex)
}
