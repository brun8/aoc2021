package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
  f, _ := os.Open("input_big.txt")
  scanner := bufio.NewScanner(f)

  input := ""

  for scanner.Scan() {
    input = scanner.Text()
  }

  s := strings.Split(input, ",")

  crabs := make([]int, 0)

  max := 0
  for _, c := range s {
    i, _ := strconv.Atoi(c)

    if i > max {
      max = i
    }
    crabs = append(crabs, i)
  }

  min := 2 << 40

  for i := 0; i < max+1; i++ {
    sum := 0
    for _, crab := range crabs {

      diff := int(math.Abs(float64(i - crab)))

      sum += GetSumTo(diff)
    }
    if sum < min {
      min = sum
    }
  }

  fmt.Println(min)
}

func GetSumTo(x int) int {
  total := 0
  for i := 1; i<x+1; i++ {
    total += i
  }
  return total
}

