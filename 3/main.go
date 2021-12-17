package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := make([]map[rune]int, 12)

	input := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()

		for i, c := range s {
			if count[i] == nil {
				count[i] = map[rune]int{}
			}
			count[i][c]++
		}
		input = append(input, s)
	}

	oxigen := filter(input, 0, true)
	co := filter(input, 0, false)

	oxigenRating, _ := strconv.ParseInt(oxigen[0], 2, 64)
	coRating, _ := strconv.ParseInt(co[0], 2, 64)

	fmt.Println(oxigenRating * coRating)
}

func filter(list []string, index int, mostCommon bool) []string {
	if len(list) == 1 {
		return list
	}

	one := []string{}
	zero := []string{}

	for _, value := range list {
		if value[index] == '1' {
			one = append(one, value)
		} else {
			zero = append(zero, value)
		}
	}

	if mostCommon {
		if len(zero) > len(one) {
			return filter(zero, index+1, mostCommon)
		} else {
			return filter(one, index+1, mostCommon)
		}
	} else {
		if len(zero) > len(one) {
			return filter(one, index+1, mostCommon)
		} else {
			return filter(zero, index+1, mostCommon)
		}
	}
}
