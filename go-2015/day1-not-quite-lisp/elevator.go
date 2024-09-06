package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// floor := 1

	data, err := os.ReadFile("./input.txt")
	check(err)

	floor := 0

	for i, char := range string(data) {
		if string(char) == ")" {
			floor = floor - 1
		} else if string(char) == "(" {
			floor = floor + 1
		}

		if floor == -1 {
			println("We headed to the basement in char: ", i+1)
			break
		}
	}
	println("The floor numer is: ", floor)
}
