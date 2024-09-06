package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	total_wrapping := 0
	total_ribbon := 0

	for _, line := range fileLines {
		dimentions_strings := strings.Split(line, "x")
		dimentions := make([]int, len(dimentions_strings))

		for i, str := range dimentions_strings {
			num, _ := strconv.Atoi(str)
			dimentions[i] = num
		}

		sort.Ints(dimentions)

		length := dimentions[0]
		width := dimentions[1]
		height := dimentions[2]

		face_three := width * length
		face_two := height * width
		face_one := length * height

		ribbon := length*2 + width*2 + (length * width * height)

		wrapping := 2 * (face_one + face_two + face_three)

		total_wrapping = total_wrapping + wrapping + length

		total_ribbon = total_ribbon + ribbon
	}
	fmt.Println(total_wrapping)
	fmt.Println(total_ribbon)
}
