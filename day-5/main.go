package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	// partTwo()
}

func partOne() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := 0

	seeds := make([]int, 0)
	seedToSoil := make(map[int]int)
	currentMap := ""

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Split(line, ": ")[0] == "seeds" {
			for _, seed := range strings.Split(strings.Split(line, ": ")[1], " ") {
				num, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, num)
			}
		}
		headerLabel := strings.Split(line, ":")
		if len(headerLabel) == 2 {
			currentMap = headerLabel[0]
			continue
		}
		nums := make([]int, 0)
		for _, val := range strings.Split(line, " ") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		if currentMap == "seed-to-soil map" {
			for i := 0; i < nums[2]; i++ {
				seedToSoil[nums[1]+i] = nums[0] + i
			}
		}
	}

	println(total)
}

// func partTwo() {
// 	f, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	total := 0

// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 	}
// }
