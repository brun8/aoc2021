package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  arr := make([]int, 0)

  for scanner.Scan() {
    line, _ := strconv.Atoi(scanner.Text())
    arr = append(arr, line)
  }

  total := 0

  for i := 1; i < len(arr)-2; i++ {
    a := arr[i-1] + arr[i] + arr[i+1]
    b := arr[i] + arr[i+1] + arr[i+2]

    if (b > a) {
      total++
    }
  }

  fmt.Println(total)
}

