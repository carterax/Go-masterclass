package main

import "fmt"

func main() {
	//var colors = make(map[string]string)
	// var colors = make(map[int]string)
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}
	// colors[10] = "#ffffff"
	//delete(colors, 10)

	// fmt.Println(colors[10])
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
