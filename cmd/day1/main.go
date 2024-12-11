package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var filepath string = "./test/data/input/day1.txt"

	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var left []int32
	var right []int32
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) > 1 {

			leftVal, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}
			left = append(left, int32(leftVal))

			rightVal, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}

			right = append(right, int32(rightVal))
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	// var freqMap map[int]int = make(map[int]int)
	// for _, leftVal := range left {
	// 	valTotal := 0
	// 	for _, rightVal := range right {
	// 		if leftVal == rightVal {
	// 			valTotal += 1
	// 		}
	// 	}
	// 	if valTotal > 0 {
	// 		freqMap[int(leftVal)] = valTotal
	// 	}
	// }
	//
	//    var total int32
	//    for k := range freqMap {
	//        total += int32(k * freqMap[k])
	//    }

	// fmt.Println(total)

	// for _, idk := range freqMap {
	//     fmt.Println(idk)
	// }
	var total int32
	for i := range len(left) {
		if left[i] > right[i] {
			diff := left[i] - right[i]
			total += int32(diff)
		} else {
			diff := right[i] - left[i]
			total += int32(diff)
		}
	}

    fmt.Println(total)

}
