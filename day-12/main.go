package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	springs := make([]Input, 0)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		numbers := make([]int, 0)
		for _, val := range strings.Split(split[1], ",") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}
		springs = append(springs, Input{split[0], numbers})
	}

	for _, input := range springs {
		sets := make([]string, 0)
		for _, str := range strings.Split(input.springs, ".") {
			if str != "" {
				sets = append(sets, str)
			}
		}
		// total := 0
		// for _, set := range sets {

		// }
	}

	// fmt.Println(springs)

}

func partTwo() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

type Input struct {
	springs string
	numbers []int
}
