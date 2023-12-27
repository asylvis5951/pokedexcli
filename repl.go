package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	nextUrl *string
	prevUrl *string
}

func printRepl() {
	fmt.Print("pokedexcli> ")
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
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
			callback:    commandMap,
		},
		// "mapb": {
		// 	name:        "mapb",
		// 	description: "Displays previous 20 location areas in Pokemon world",
		// 	callback:    commandMapb,
		// },
	}
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	printRepl()
	for scanner.Scan() {
		i := scanner.Text()
		if c, ok := getCommands()[i]; ok {
			c.callback()
		}
		printRepl()
	}
}
