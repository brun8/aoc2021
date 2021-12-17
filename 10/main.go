package main

import (
	"bufio"
	"fmt"
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

  dict := make(map[byte]byte)
  dict[byte(']')] = '['
  dict[byte('}')] = '{'
  dict[byte(')')] = '('
  dict[byte('>')] = '<'

  //stack := make([]byte, 0)
  stack := []byte{byte('['), byte(']')}
  fmt.Println(stack)
  Push('a', &stack)
  fmt.Println(stack)
  Pop(&stack)
  fmt.Println(stack)
  //for _, line := range input {
    //for _, c := range line {
      //if dict[byte(c)] == 0 {
      //} else {
      //}
    //}
  //}
}

func Push(c byte, s *[]byte) {
  *s = append(*s, c)
}

func Pop(s *[]byte) byte {
  sr := *s
  c := sr[len(sr)-1]
  return c
}
