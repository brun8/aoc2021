package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  input := parseInput()

  numbers := strings.Split(input[0], ",")
  input = input[1:]
  boards := parseBoards(input)
  drawed := make(map[string]bool)

  fmt.Println(numbers)
  won := make(map[int]bool)
  total := 0

  for i := 0; i < len(numbers); i++ {
    drawed[numbers[i]] = true
    last := numbers[i]

    for j, board := range boards {
      if won[j] == true {
        fmt.Printf("board %d won, skipping\n", j)
      } else {
        if checkBoard(board, drawed) {
          won[j] = true
          total++

          if total == len(boards) {
            printBoard(board)
            fmt.Println("last: ", last)
            fmt.Println(calculatePoints(board, last, drawed))
            return
          }
        }
      }
    }
  }
}

func calculatePoints(board []string, last string, draws map[string]bool) int {
  parsed := [][]string{}
  for i := range board {
    line := strings.Fields(board[i])
    parsed = append(parsed, line)
  }

  lastNum, _ := strconv.Atoi(last)

  sum := 0
  for _, line := range parsed {
    for _, cell := range line {
      num, _ := strconv.Atoi(cell)
      if !draws[cell] {
        sum += num
      }
    }
  }

  total := sum * lastNum


  return total
}

func checkBoard(board []string, draws map[string]bool) bool {
  parsed := [][]string{}

  for i := range board {
    line := strings.Fields(board[i])
    parsed = append(parsed, line)
  }

  for i := 0; i<5; i++ {
    vertical := true
    horizontal := true
    mdiagonal := false
    sdiagonal := false

    for j := 0; j<5; j++ {
      currenth := parsed[i][j]
      currentv := parsed[j][i]

      if !draws[currenth] {
        horizontal = false
      }

      if !draws[currentv] {
        vertical = false
      }
    }

    if vertical || horizontal || mdiagonal || sdiagonal {
      return true
    }
  }
  return false
}

func printBoard(b []string) {
  for _, line := range b {
    fmt.Println(line)
  }
}

func parseBoards(input []string) [][]string {
  boards := make([][]string, 0)
  for i := 0; i < len(input); i+=5 {
    board := make([]string, 0)
    for j := 0; j < 5; j++ {
      board = append(board, input[i+j])
    }
    boards = append(boards, board)
  }

  return boards
}

func parseInput() []string {
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  input := []string{}

  for scanner.Scan() {
    line := scanner.Text()

    if len(line) > 0 {
      input = append(input, scanner.Text())
    }
  }

  return input
}
