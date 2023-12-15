package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputs := ""

	for scanner.Scan() {
		inputs += scanner.Text()
	}

	currentValue := 0

	for _, input := range strings.Split(inputs, ",") {
		currentValue += HASH(input)
	}

	println(currentValue)

}

func partTwo() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// input := ""

	// for scanner.Scan() {
	// 	input += scanner.Text()
	// }
}

func HASH(input string) int {
	currentValue := 0
	for _, ascii := range input {
		currentValue += int(ascii)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
