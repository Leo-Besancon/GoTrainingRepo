package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	printMap(colors)

}

func printMap(mapping map[string]string) {
	for color, hex := range mapping {
		fmt.Println(color, " : ", hex)
	}
}
