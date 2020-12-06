package main

import (
	"bufio"
	"log"
	"os"
    "strings"
)

type Empty struct{}
type Set map[string]Empty

func NewSet(xs []string) Set {
    var e Empty
    s := make(Set)
    for _, x := range xs {
        s[x] = e
    }
    return s
}

func (s1 Set) isSubset(s2 Set) bool {
    for k, _ := range s1 {
        if _, ok := s2[k]; !ok {
            return false
        }
    }
    return true
}

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
    spec := NewSet([]string{
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    })
    fields := make([]string, 0)
    validCount := 0
    nextPassport := func() {
        if spec.isSubset(NewSet(fields)) {
            validCount++
        }
        fields = fields[:0]
    }

	for {
        if !scanner.Scan() {
            nextPassport()
            break
        }

		rawFields := strings.Fields(scanner.Text())
        if len(rawFields) == 0 {
            nextPassport()
            continue
        }

        for _, rawField := range rawFields {
            field := strings.Split(rawField, ":")
            fields = append(fields, field[0])
        }
	}

    log.Println(validCount)
}
