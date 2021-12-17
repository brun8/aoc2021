package main

import (
	"bufio"
	"os"
)

func main() {
  f, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(f)
  defer f.Close()

  input := make([]string, 0)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
}

