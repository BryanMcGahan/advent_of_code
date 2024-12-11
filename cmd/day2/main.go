package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const FILE string = "./test/data/input/day2.txt"

func tryLine(line []int32) bool {
	if safeLine(line) {
		return true
	}

	for i := 0; i < len(line); i++ {
		var copiedLine []int32 = make([]int32, len(line))
		copy(copiedLine, line)
		copiedLine = append(copiedLine[:i], copiedLine[i+1:]...)
        if safeLine(copiedLine) {
            return true
        }
	}

	return false
}

func safeLine(line []int32) bool {
	var inc bool = line[0] < line[1]
	for i := 1; i < len(line); i++ {
		diff := int(math.Abs(float64(line[i] - line[i-1])))
		if diff < 1 || diff > 3 || (inc && line[i] < line[i-1]) || (!inc && line[i] > line[i-1]) {
			return false
		}
	}

	return true
}

func main() {

	file, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(file)
	var strLines []string = strings.Split(contents, "\n")
	var lines [][]int32 = make([][]int32, len(strLines)-1)

	for j, line := range strLines {
		chars := strings.Split(line, " ")
		numLine := make([]int32, len(chars))

		if len(chars) <= 1 {
			continue
		}

		for i := 0; i < len(chars); i++ {
			num, err := strconv.Atoi(string(chars[i]))
			if err != nil {
				log.Fatal(err)
			}

			numLine[i] = int32(num)
		}
		lines[j] = numLine
	}

	var safe int = 0
	for _, line := range lines {
		isSafe := tryLine(line)
		if isSafe {
			safe++
		}
	}

	fmt.Println("SAFE: ", safe)
}
