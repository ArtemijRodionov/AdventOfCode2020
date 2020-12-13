package main

import (
	"bufio"
	"log"
	"os"
)

func parse(s string, lo, hi int, loChar, hiChar rune) int {
	for _, c := range s[:len(s)-1] {
		switch c {
		case loChar:
			hi -= (hi - lo + 1) / 2
		case hiChar:
			lo += (hi - lo + 1) / 2
		}
	}
	if rune(s[len(s)-1]) == hiChar {
		return hi
	}
	return lo
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxId := 0
	idPass := make(map[int]struct{})
	for scanner.Scan() {
		raw := scanner.Text()
		row := parse(raw[:7], 0, 127, 'F', 'B')
		col := parse(raw[7:], 0, 7, 'L', 'R')
		id := row*8 + col
		idPass[id] = struct{}{}
		if maxId < id {
			maxId = id
		}
	}
	missedId := 0
	offsets := []int{1, -1}
	for id := range idPass {
		for x := range offsets {
			_, boundOk := idPass[id+2*x]
			_, centerOk := idPass[id+1*x]
			if boundOk && !centerOk {
				missedId = id + 1*x
				break
			}
		}
	}
	log.Println(maxId)
	log.Println(missedId)
}
