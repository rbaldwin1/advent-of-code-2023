package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	partOne()
}

func partOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0

	for i, line := range lines {
		currentNumber := ""
		for j, c := range line {
			char := byte(c)
			if isNumber(char) {
				currentNumber += string(char)
			}
			if currentNumber != "" && (!isNumber(char) || j == len(line)-1) {
				start := j - (len(currentNumber) + 1)
				if j == len(line)-1 && isNumber(char) {
					start = j - len(currentNumber)
				}
				if start < 0 {
					start = 0
				}
				lineIndicesToCheck := make([]int, 0)
				if i > 0 {
					lineIndicesToCheck = append(lineIndicesToCheck, i-1)
				}
				lineIndicesToCheck = append(lineIndicesToCheck, i)
				if i < len(lines)-1 {
					lineIndicesToCheck = append(lineIndicesToCheck, i+1)
				}
				symbolFound := false
				for _, lineIndexToCheck := range lineIndicesToCheck {
					if symbolFound {
						break
					}
					lineToCheck := lines[lineIndexToCheck]
					for k := start; k <= j; k++ {
						if isSymbol(lineToCheck[k]) {
							symbolFound = true
							break
						}
					}
				}

				if symbolFound {
					num, err := strconv.Atoi(currentNumber)
					if err != nil {
						panic(err)
					}
					total += num
				}

				currentNumber = ""
			}
		}
	}

	println(total)

}

func isNumber(value byte) bool {
	return value >= 48 && value <= 57
}

func isSymbol(value byte) bool {
	return !isNumber(value) && value != 46
}
