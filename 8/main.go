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
  inputs := make([][]string, 0)
  outputs := make([][]string, 0)


  for scanner.Scan() {
    split := strings.Split(scanner.Text(), " | ")

    curInput := strings.Fields(split[0])
    curOutput := strings.Fields(split[1])

    inputs = append(inputs, curInput)
    outputs = append(outputs, curOutput)
  }

  total := 0
  for i, in := range inputs {
    // mapa de numeros ja descobertos
    decoded := make(map[string]string)

    // mapa de len do numero para numeros codificados do tamanho da chave
    lenMap := make(map[int][]string)

    // mapa de letras ja descobertas
    dict := make(map[string]string)

    for _, entry := range in {
      l := len(entry)

      lenMap[l] = append(lenMap[l], entry)
    }

    decoded["1"] = lenMap[2][0]
    decoded["7"] = lenMap[3][0]
    decoded["4"] = lenMap[4][0]
    decoded["8"] = lenMap[7][0]

    for _, s := range lenMap[6] {
      diff := AnotInB(decoded["1"], s)

      if len(diff) != 0 {
        // acha o numero 6 e aresta 'c'
        decoded["6"] = s
        dict["c"] = diff[0]
      }
    }

    // acha 3 comparando com 1
    for _, s := range lenMap[5] {
      r := AnotInB(decoded["1"], s)

      if len(r) == 0 {
        decoded["3"] = s
      }
    }

    // acha 2 e 5 comparando com 1
    for _, s := range lenMap[5] {
      r := AnotInB(decoded["1"], s)

      if len(r) == 1 {
        if r[0] == dict["c"] {
          decoded["5"] = s
        } else {
          decoded["2"] = s
        }
      }
    }

    // acha 9 e 0 comparando com 4
    for _, s := range lenMap[6] {
      diff := AnotInB(decoded["4"], s)

      if len(diff) == 0 {
        decoded["9"] = s
      } else {
        if diff[0] != dict["c"] {
          decoded["0"] = s
          dict["d"] = diff[0]
        }
      }
    }

    out := outputs[i]

    result := ""
    for _, o := range out {
      for i := 0; i < 10; i++ {
        s := fmt.Sprint(i)
        cur := decoded[s]
        // gambiarra pq as letras nao estao sempre na mesma ordem
        // tem como melhorar muito
        if len(o) == len(cur) {
          diff := GetDiff(o, cur)
          if len(diff) == 0 {
            result += s
          }
        }
      }
    }
    parsed, _ := strconv.Atoi(result)
    total += parsed
  }
  fmt.Println(total)
}

// retorna slice de letras que sao diferentes
func GetDiff(a, b string) []string {
  result := make([]string, 0)

  countA := make(map[string]int)
  countB := make(map[string]int)

  for _, c := range a {
    countA[string(c)]++
  }

  for _, c := range b {
    countB[string(c)]++
  }

  for _, c := range "abcdefg" {
    if countA[string(c)] != countB[string(c)] {
      result = append(result, string(c))
    }
  }
  return result
}

// retorna slice de letras que estao em A e nao em B
func AnotInB(a, b string) []string {
  result := make([]string, 0)

  countA := make(map[rune]int)
  countB := make(map[rune]int)

  for _, c := range a {
    countA[c]++
  }
  for _, c := range b {
    countB[c]++
  }

  for _, c := range "abcdefg" {
    if countA[c] == 1 && countB[c] == 0 {
      result = append(result, string(c))
    }
  }
  return result
}

