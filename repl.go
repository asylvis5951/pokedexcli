package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/asylvis5951/pokedexcli/internal/pokecache"
	pokedex "github.com/asylvis5951/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, *pokedex.Pdex, string) error
}

type config struct {
	nextUrl *string
	prevUrl *string
}

func printRepl() {
	fmt.Print("pokedexcli> ")
}

func commandExit(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, s string) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, cache *pokecache.Cache, p *pokedex.Pdex, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n\n", cmd.name, cmd.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays names of 20 location areas in Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas in Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays information about a location-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch specified pokemon.  Ex. catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:	"inspect",
			description: "Displays entry for pokemon in Pokedex, if it exists",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays pokemon in pokedex",
			callback: commandPdex,
		},
	}
}

func startRepl(cfg *config, cache *pokecache.Cache, pokedex *pokedex.Pdex) {
	scanner := bufio.NewScanner(os.Stdin)
	printRepl()
	for scanner.Scan() {
		i := scanner.Text()
		slc := strings.Split(i, " ")
		if len(slc) > 1 {
			if c, ok := getCommands()[slc[0]]; ok {
				c.callback(cfg, cache, pokedex, slc[1])
			}
		} else {
			if c, ok := getCommands()[slc[0]]; ok {
				c.callback(cfg, cache, pokedex, slc[0])
			}
		}
		printRepl()
	}
}
