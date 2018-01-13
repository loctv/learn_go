package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "red")
	fmt.Println(colors)
}
