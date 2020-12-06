package main

import (
	"bufio"
	"log"
	"os"
)

type Slope struct{ X, Y int }

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	slopes := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}

	slopeTreeCount := make(map[Slope]int)
	for _, slope := range slopes {
		for i, row := 1, slope.Y; row < len(grid); i, row = i+1, row+slope.Y {
			col := (i * slope.X) % len(grid[row])
			if rune(grid[row][col]) == '#' {
				slopeTreeCount[slope]++
			}
		}
	}
	log.Println(slopeTreeCount)
	result := 1
	for _, count := range slopeTreeCount {
		result *= count
	}
	log.Println(result)
}
