package main

import (
	"bufio"
	"math"
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

	universe := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if allSame(line, '.') {
			newLine := strings.Clone(line)
			universe = append(universe, newLine)
		}
		universe = append(universe, line)
	}

	i := 0
	max := len(universe[0])
	for i < max {
		allSame := true
		for _, str := range universe {
			if str[i] != '.' {
				allSame = false
				break
			}
		}
		if allSame {
			for row, str := range universe {
				universe[row] = str[:i] + "." + str[i:]
			}
			max = len(universe[0])
			i++
		}
		i++
	}

	galaxies := make([]Coord, 0)
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == '#' {
				galaxies = append(galaxies, Coord{row, col})
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += int(math.Abs(float64(galaxies[i].row - galaxies[j].row)))
			total += int(math.Abs(float64(galaxies[i].col - galaxies[j].col)))
		}
	}

	println(total)

}

func partTwo() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// sum := 0

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	nums := strToIntArray(line)
	// 	prev := getPrevValue(nums)
	// 	sum += prev
	// }

	// println(sum)
}

func allSame(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != char {
			return false
		}
	}
	return true
}

type Pair struct {
	first  Coord
	second Coord
}

type Coord struct {
	row int
	col int
}
