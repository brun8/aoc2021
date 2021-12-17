package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  arr := make([]string, 0)

  for scanner.Scan() {
    arr = append(arr, scanner.Text())
  }

  depth := 0
  horizontal := 0
  aim := 0

  for _, line := range arr {
    words := strings.Fields(line)

    times, _ := strconv.Atoi(words[1])

    if words[0] == "down" {
      aim += times
    } else if words[0] == "up" {
      aim -= times
    } else if words[0] == "forward" {
      horizontal += times
      depth += aim * times
    } else {
      horizontal -= times
    }
  }
  fmt.Println(horizontal*depth)
}

