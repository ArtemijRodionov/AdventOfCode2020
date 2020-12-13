package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Validator func(string) bool

func inRange(from, to int) Validator {
	return func(s string) bool {
		if val, err := strconv.Atoi(s); err == nil {
			return from <= val && to >= val
		}
		return false
	}
}
func hasPattern(exp string) Validator {
	re := regexp.MustCompile(exp)
	return func(s string) bool {
		return re.Match([]byte(s))
	}
}
func const_(x bool) Validator {
	return func(s string) bool { return x }
}
func or(xs ...Validator) Validator {
	return func(s string) bool {
		for _, x := range xs {
			if x(s) {
				return true
			}
		}
		return false
	}
}

var validators = map[string]Validator{
	"byr": inRange(1920, 2002),
	"iyr": inRange(2010, 2020),
	"eyr": inRange(2020, 2030),
	"hgt": or(
		hasPattern("(1[5-8][0-9]|19[0-3])cm"),
		hasPattern("(59|6[0-9]|7[0-6])in"),
	),
	"hcl": hasPattern("#[0-9a-f]{6}"),
	"ecl": hasPattern("(amb|blu|brn|gry|grn|hzl|oth)"),
	"pid": hasPattern("[0-9]{9}"),
	"cid": const_(true),
}

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

func main() {
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

	validCount := 0
	validCount2 := 0
	fieldVals := make(map[string]string, 0)

	nextPassport := func() {
		isValid := true
		fields := make([]string, len(fieldVals))
		for f, v := range fieldVals {
			fields = append(fields, f)
			if validator, ok := validators[f]; !(ok && validator(v)) {
				isValid = false
			}
		}
		if spec.isSubset(NewSet(fields)) {
			validCount++
		} else {
			isValid = false
		}
		if isValid {
			validCount2++
		}
		fieldVals = make(map[string]string)
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
			parsed := strings.Split(rawField, ":")
			field, val := parsed[0], parsed[1]
			fieldVals[field] = val
		}
	}

	log.Println(validCount)
	log.Println(validCount2 - 1)
}
