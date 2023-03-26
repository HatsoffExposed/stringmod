package main

import (
	agame "chpt13/geometry"
	geometry "chpt13/geometry/coordinate"
	"fmt"

	"github.com/hackebrot/turtle"
)

func main() {
	pt1 := geometry.Point{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Length())
	agame.Ex()

	emoji, ok := turtle.Emojis["smiley"]

	if !ok {
		fmt.Println("No emoji found")
	} else {
		fmt.Println(emoji.Char)
	}
}
