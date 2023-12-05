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
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	seeds := make([]int, 0)
	currentMaps := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Split(line, ": ")[0] == "seeds" {
			for _, seed := range strings.Split(strings.Split(line, ": ")[1], " ") {
				num, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, num)
			}
			continue
		}
		headerLabel := strings.Split(line, ":")
		if len(headerLabel) == 2 {
			currentMaps = make([][]int, 0)
			continue
		}
		if line == "" {
			// Transform values via current map
			for i, seed := range seeds {
				for _, mapn := range currentMaps {
					if seed >= mapn[1] && seed < mapn[1]+mapn[2] {
						seeds[i] = mapn[0] + (seed - mapn[1])
					}
				}
			}
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
		currentMaps = append(currentMaps, nums)
	}

	// Final transform after file ends
	answer := seeds[0]
	for i, seed := range seeds {
		for _, mapn := range currentMaps {
			if seed >= mapn[1] && seed < mapn[1]+mapn[2] {
				seeds[i] = mapn[0] + (seed - mapn[1])
			}
		}
		if seeds[i] < answer {
			answer = seeds[i]
		}
	}

	println(answer)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	seeds := make([]int, 0)
	currentMaps := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		seedRanges := make([]int, 0)
		if strings.Split(line, ": ")[0] == "seeds" {
			for _, seed := range strings.Split(strings.Split(line, ": ")[1], " ") {
				num, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seedRanges = append(seedRanges, num)
			}
			for i := 0; i < len(seedRanges); i++ {
				if i%2 != 0 {
					continue
				}
				start := seedRanges[i]
				end := start + seedRanges[i+1]
				for j := start; j < end; j++ {
					seeds = append(seeds, j)
				}
			}
			continue
		}
		headerLabel := strings.Split(line, ":")
		if len(headerLabel) == 2 {
			currentMaps = make([][]int, 0)
			continue
		}
		if line == "" {
			// Transform values via current maps
			for i, seed := range seeds {
				for _, mapn := range currentMaps {
					if seed >= mapn[1] && seed < mapn[1]+mapn[2] {
						seeds[i] = mapn[0] + (seed - mapn[1])
					}
				}
			}
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
		currentMaps = append(currentMaps, nums)
	}

	// Final transform after file ends
	answer := seeds[0]
	for i, seed := range seeds {
		for _, mapn := range currentMaps {
			if seed >= mapn[1] && seed < mapn[1]+mapn[2] {
				seeds[i] = mapn[0] + (seed - mapn[1])
			}
		}
		if seeds[i] < answer {
			answer = seeds[i]
		}
	}

	println(answer)

}
