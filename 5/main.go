package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  //"math"
	"strings"
)

type Point struct {
  x int
  y int
}

type Vent struct {
  start Point
  end Point
}

func main() {
  f, _ := os.Open("input_big.txt")
  scanner := bufio.NewScanner(f)
  input := make([]string, 0)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }

  vents, max_x, max_y := ParseLines(input)

  points := make([]Point, 0)

  for _, v := range vents {
    result := GetPointsBetween(v)

    for _, p := range result {
      points = append(points, p)
    }
  }

  diagram := make(map[Point]int)
  total := 0

  for _, p := range points {
    if diagram[p] == 1 {
      diagram[p]++
      total++
    } else {
      diagram[p]++
    }
  }

  for i := 0; i < max_x+1; i++ {
    s := ""
    for j := 0; j < max_y+1; j++ {
      p := Point{j, i}
      if diagram[p] == 0 {
        s += "."
      } else {
        s += fmt.Sprint(diagram[p])
      }
    }
    fmt.Println(s)
  }

  fmt.Println(total)

}

func GetPointsBetween(vent Vent) []Point {
  points := make([]Point, 0)

  start := vent.start
  end := vent.end

  orientation := vent.GetVentType()

  var cur int

  if orientation == "x" {
    cur = start.x

    for cur != end.x {
      points = append(points, Point{cur, start.y})

      if cur > end.x {
        cur--
      } else {
        cur++
      }
    }
    //points = append(points, Point{end.x, end.y})
  } else if orientation == "y" {
    cur = start.y
    for cur != end.y {
      points = append(points, Point{start.x, cur})
      if cur > end.y {
        cur--
      } else {
        cur++
      }
    }
    //points = append(points, Point{end.x, end.y})
  } else {
    curx := start.x
    cury := start.y

    for curx != end.x {
      points = append(points, Point{curx, cury})

      if curx > end.x {
        curx--
      } else {
        curx++
      }

      if cury > end.y {
        cury--
      } else {
        cury++
      }
    }
  }
  points = append(points, Point{end.x, end.y})
  return points
}

func (vent Vent) GetVentType() string {
  if vent.start.x == vent.end.x {
    return "y"
  } else if vent.start.y == vent.end.y {
    return "x"
  }
  return "d"
}

func (v Vent) IsStraight() bool {
  return v.start.x == v.end.x || v.start.y == v.end.y
}

func ParseLines(input []string) ([]Vent, int, int) {
  vents := make([]Vent, 0)
  var max_x, max_y int

  for _, line := range input {
    points := strings.Split(line, " -> ")

    start := strings.Split(points[0], ",")
    end := strings.Split(points[1], ",")

    x1, _ := strconv.Atoi(start[0])
    y1, _ := strconv.Atoi(start[1])

    if x1 > max_x {
      max_x = x1
    }

    if y1 > max_y {
      max_y = y1
    }

    x2, _ := strconv.Atoi(end[0])
    y2, _ := strconv.Atoi(end[1])

    if x2 > max_x {
      max_x = x2
    }

    if y2 > max_y {
      max_y = y2
    }

    startPoint := Point{x1, y1}
    endPoint := Point{x2, y2}

    vent := Vent{start: startPoint, end: endPoint}

    vents = append(vents, vent)
  }
  return vents, max_x, max_y
}


