package main

import (
	"fmt"
	"github.com/asylvis5951/pokeapi"
)

func commandMapf() error {
	// Get current state
	e := "location-area/"
	next, prev, err := pokeapi.Get(e)
	fmt.Println(next)
	fmt.Println(prev)
	if err != nil {
		return err
	}
	return nil
}
