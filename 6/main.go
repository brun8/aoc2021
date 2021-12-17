package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f, _ := os.Open("input_big.txt")
  scanner := bufio.NewScanner(f)

  var input string

  for scanner.Scan() {
    input = scanner.Text()
  }

  s := strings.Split(input, ",")

  fish := make(map[int]int, 0)

  for _, f := range s {
    cur, _ := strconv.Atoi(f)
    fish[cur]++
  }

  for i := 0; i < 256; i++ {
    prev := fish[0]
    for j := 8; j >= 0; j-- {
      cur := fish[j]
      if j == 0 {
        fish[6] += cur
        fish[j] = prev
      }
      fish[j] = prev
      prev = cur
    }
  }

  total := 0

  for i := 0; i<=8; i++ {
    total += fish[i]
  }

  fmt.Println(fish)
  fmt.Println(total)
}

