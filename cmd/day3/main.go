package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INPUT string = "./test/data/input/day3.txt"

// const INPUT string = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func main() {

	data, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}

	inputString := string(data)

	doDontRegEx, err := regexp.Compile(`do\(\)|don\'t\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	mulRegEx, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		log.Fatal(err)
	}

	foundMuls := mulRegEx.FindAllString(inputString, -1)
	foundMulsIdxs := mulRegEx.FindAllStringSubmatchIndex(inputString, -1)
    fmt.Println(len(foundMuls))

	foundDoOrDont := doDontRegEx.FindAllString(inputString, -1)
	foundDoOrDontIdxs := doDontRegEx.FindAllStringSubmatchIndex(inputString, -1)
    fmt.Println(len(foundDoOrDont))

	doOrDontMap := make(map[bool][][]int)
	previousFound := 0
	prevDo := true
	for i, doOrDont := range foundDoOrDont {
		doOrDontIdxStart := foundDoOrDontIdxs[i][0]
		// doOrDontIdxEnd := foundDoOrDontIdxs[i][1]
		var rng [2]int = [2]int{previousFound, doOrDontIdxStart}

		if doOrDont == "do()" {
			if !prevDo {
				doOrDontMap[false] = append(doOrDontMap[false], rng[:])
				prevDo = true
				previousFound = doOrDontIdxStart
			}

		} else if doOrDont == "don't()" {
			if prevDo {
				doOrDontMap[true] = append(doOrDontMap[true], rng[:])
				prevDo = false
				previousFound = doOrDontIdxStart
			}

		}
	}

	if prevDo {
		var rng [2]int = [2]int{previousFound, len(inputString)}
		doOrDontMap[true] = append(doOrDontMap[true], rng[:])
	} else {
		var rng [2]int = [2]int{previousFound, len(inputString)}
		doOrDontMap[false] = append(doOrDontMap[false], rng[:])
	}

	var total int64 = 0
	for i, res := range foundMuls {
		mulsIdx := foundMulsIdxs[i][0]
		for _, rng := range doOrDontMap[true] {
			if mulsIdx >= rng[0] && mulsIdx <= rng[1] {

				parts := strings.Split(res, ",")

				left := parts[0]
				leftParts := strings.Split(left, "(")
				leftNum, err := strconv.Atoi(leftParts[1])
				if err != nil {
					log.Fatal(err)
				}

				right := parts[1]
				rightParts := strings.Split(right, ")")
				rightNum, err := strconv.Atoi(rightParts[0])
				if err != nil {
					log.Fatal(nil)
				}

				total += int64(int64(leftNum) * int64(rightNum))
				break
			}
		}

	}

	fmt.Println(total)

    var allTotal int64 = 0
	for _, res := range foundMuls {

		parts := strings.Split(res, ",")

		left := parts[0]
		leftParts := strings.Split(left, "(")
		leftNum, err := strconv.Atoi(leftParts[1])
		if err != nil {
			log.Fatal(err)
		}

		right := parts[1]
		rightParts := strings.Split(right, ")")
		rightNum, err := strconv.Atoi(rightParts[0])
		if err != nil {
			log.Fatal(nil)
		}

		allTotal += int64(int64(leftNum) * int64(rightNum))
	}

    fmt.Println(allTotal)
}
