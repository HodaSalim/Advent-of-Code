package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	directions, err := os.ReadFile("./input.txt")
	check(err)

	visited := make(map[[2]int]bool)
	santaX, santaY := 0, 0
	roboX, roboY := 0, 0
	visited[[2]int{santaX, santaY}] = true

	for i, direction := range directions {
		if i%2 == 0 {
			// Santa's turn
			switch direction {
			case '^':
				santaY++
			case 'v':
				santaY--
			case '>':
				santaX++
			case '<':
				santaX--
			}
			visited[[2]int{santaX, santaY}] = true
		} else {
			// Robo-Santa's turn
			switch direction {
			case '^':
				roboY++
			case 'v':
				roboY--
			case '>':
				roboX++
			case '<':
				roboX--
			}
			visited[[2]int{roboX, roboY}] = true
		}
	}
	fmt.Println(len(visited))
}
