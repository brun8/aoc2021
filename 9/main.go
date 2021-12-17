package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
  i, j int
}

func main() {
  f, _ := os.Open("input_big.txt")
  scanner := bufio.NewScanner(f)
  defer f.Close()

  input := make([][]int, 0)
  for scanner.Scan() {
    parsed := make([]int, 0)
    numbers := strings.Split(scanner.Text(), "")
    for _, n := range numbers {
      converted, _ := strconv.Atoi(n)
      parsed = append(parsed, converted)
    }
    input = append(input, parsed)
  }

  height := len(input)
  width := len(input[0])

  basins := make([]int, 0)

  for i := 0; i < height; i++ {
    for j := 0; j < width; j++ {
      b := FindBasin(i, j, input)
      if b != 0 {
        basins = append(basins, b)
      }
    }
  }
  sort.Slice(basins, func (a, b int) bool { return basins[a] > basins[b] })
  t := 1
  for i := 0; i < 3; i++ {
    t *= basins[i]
  }
  fmt.Println(t)

}

func InBounds(i, j int, m [][]int) bool {
  return i >= 0 && i <= len(m)-1 && j >= 0 && j <= len(m[i])-1
}

func FindBasin(i, j int, m [][]int) int {
  if !InBounds(i, j, m) {
    return 0
  }

  //fmt.Println(m)
  cur := m[i][j]
  if cur == 9 {
    return 0
  }

  m[i][j] = 9
  dirs := []Dir{{-1,0},{1,0},{0,-1},{0,1}}
  t := 0
  for _, d := range dirs {
    di := i + d.i
    dj := j + d.j
    t += FindBasin(di, dj, m)
  }
  //fmt.Println(i, j, t)
  return t + 1
}
