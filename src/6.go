package main

import (
	"bufio"
	"log"
	"os"
)

type Empty struct{}

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
    raw := ""
    count := 0
    answers := make(map[rune]Empty)
	for {
        isOk := scanner.Scan()
        if isOk {
            raw = scanner.Text()
            for _, r := range raw {
                answers[r] = Empty{}
            }
        }

        if raw == "" || !isOk {
            count += len(answers)
            answers = make(map[rune]Empty)
        }

        if !isOk { break }
	}
    log.Println(count)
}
