package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	raw := ""
	anyCount := 0
	allCount := 0
	groupSize := 0
	yes := make(map[rune]int)
	for {
		isOk := scanner.Scan()
		if isOk {
			raw = scanner.Text()
			for _, r := range raw {
				yes[r]++
			}
		}

		if raw == "" || !isOk {
			for _, count := range yes {
				if count == groupSize {
					allCount++
				}
			}
			anyCount += len(yes)
			groupSize = 0
			yes = make(map[rune]int)
		} else {
			groupSize++
		}

		if !isOk {
			break
		}
	}
	log.Println(anyCount)
	log.Println(allCount)
}
