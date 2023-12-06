package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// partOne()
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
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// seeds := make([]int, 0)
	allMaps := make([][][]int, 0)
	currentMaps := make([][]int, 0)
	seedRanges := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentMaps) > 0 {
				allMaps = append(allMaps, currentMaps)
			}
			continue
		}
		if strings.Split(line, ": ")[0] == "seeds" {
			for _, seed := range strings.Split(strings.Split(line, ": ")[1], " ") {
				num, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seedRanges = append(seedRanges, num)
			}
			continue
		}
		headerLabel := strings.Split(line, ":")
		if len(headerLabel) == 2 {
			currentMaps = make([][]int, 0)
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

	allMaps = append(allMaps, currentMaps)

	for _, set := range allMaps {
		slices.SortFunc(set, sortMapSets)
	}

	ranges := make([]valRange, 0)

	for i := 0; i < len(seedRanges); i += 2 {
		ranges = append(ranges, valRange{seedRanges[i], seedRanges[i] + (seedRanges[i+1] - 1)})
	}

	// DEBUG - print ranges
	// for _, r := range ranges {
	// 	fmt.Println(r)
	// }

	// for _, mapn := range allMaps {
	// 	mapInputRanges := make([]valRange, 0)
	// 	for _, mapr := range mapn {

	// 	}
	// }

	// newRanges := make([]valRange, 0)
	// for _, mapn := range allMaps[0] {

	// }

	// for i := len(allMaps) - 1; i >= 0; i-- {
	// 	for j := 0; j < len(allMaps[i]); j++ {
	// 		min := allMaps[j][0]
	// 		max := min + (allMaps[j][2] - 1)
	// 	}
	// 	min := allMaps[i][0]
	// 	max := min + (allMaps[i][2] - 1)
	// 	for j := i - 1; j >= 0; j-- {

	// 	}
	// }

	// Final transform after file ends
	// answer := seeds[0]

	// println(answer)

}

func sortMapSets(mapSet1 []int, mapSet2 []int) int {
	if mapSet1[0] < mapSet2[0] {
		return -1
	}
	return 1
}

type valRange struct {
	min int
	max int
}
