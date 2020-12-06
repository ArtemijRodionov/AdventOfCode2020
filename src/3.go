package main

import (
	"bufio"
	"log"
	"os"
)

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	treeCount := 0
	for row := 1; row < len(grid); row++ {
		col := (row * 3) % len(grid[row])
		if rune(grid[row][col]) == '#' {
			treeCount++
		}
	}
	log.Println(treeCount)
}
