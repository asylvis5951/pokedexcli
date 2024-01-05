package main

import (
	"encoding/json"
	"fmt"
	"github.com/asylvis5951/pokedexcli/internal/pokeapi"
	"github.com/asylvis5951/pokedexcli/internal/pokecache"
	"github.com/asylvis5951/pokedexcli/internal/pokedex"
)

func commandInspect(cfg *config, cache *pokecache.Cache, pd *pokedex.Pdex, name string) error {
	if body, ok := pd.Get(name); ok {
		var p pokeapi.Pokemon
		json.Unmarshal(body, &p)
		fmt.Printf("Name: %v\n", p.Name)
		fmt.Printf("Height: %v\n", p.Height)
		fmt.Printf("Stats: \n")
		fmt.Printf("  - hp: %v\n", p.Stats[0].BaseStat)
		fmt.Printf("  - attack: %v\n", p.Stats[1].BaseStat)
		fmt.Printf("  - defense: %v\n", p.Stats[2].BaseStat)
		fmt.Printf("  - special-attack: %v\n", p.Stats[3].BaseStat)
		fmt.Printf("  - special-defense: %v\n", p.Stats[4].BaseStat)
		fmt.Printf("  - speed: %v\n", p.Stats[5].BaseStat)
		fmt.Printf("Types: \n")
		for _, v := range p.Types {
			fmt.Printf("  - %v\n", v.Type.Name)
		}
		return nil
	} else {
		fmt.Printf("%s does not exist in Pokedex!\n", name)
		return fmt.Errorf("%s does not exist in Pokedex!\n", name)
	}
}
