package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Counter map[string]int
type Index map[string]Counter

const (
	RuleSep = "contain"
	CapSep  = ","
	NumSep  = " "
	ToTrim  = ". "
)

var ToRemove = strings.NewReplacer(
	"bags", "",
	"bag", "",
)

func traverse(idx Index, name string, fn func(string)) {
	outBags, ok := idx[name]
	if !ok {
		return
	}
	for bagName := range outBags {
		fn(bagName)
		traverse(idx, bagName, fn)
	}
}

func countBags(idx Index, name string) int {
	bags, ok := idx[name]
	if !ok {
		return 0
	}
	result := 0
	for bag, count := range bags {
		result += count + count*countBags(idx, bag)
	}
	return result
}

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
	innerOuter := make(Index)
	outerInner := make(Index)
	for scanner.Scan() {
		line := scanner.Text()

		ruleCaps := strings.Split(line, RuleSep)
		caps := strings.Split(strings.TrimSpace(ruleCaps[1]), CapSep)
		outer := strings.TrimSpace(ToRemove.Replace(ruleCaps[0]))
		outerInner[outer] = make(Counter)
		for _, cap_ := range caps {
			numName := strings.SplitN(strings.Trim(cap_, ToTrim), NumSep, 2)
			rawCapNum := strings.TrimSpace(numName[0])
			capNum, err := strconv.Atoi(rawCapNum)
			if err != nil {
				capNum = 0
			}
			inner := strings.TrimSpace(ToRemove.Replace(numName[1]))
			outerInner[outer][inner] = capNum
			if _, ok := innerOuter[inner]; !ok {
				innerOuter[inner] = make(Counter)
			}
			innerOuter[inner][outer] = capNum
		}
	}

	bag := "shiny gold"
	outerBags := make(Counter)
	traverse(innerOuter, bag, func(bagName string) {
		outerBags[bagName]++
	})
	innerCount := countBags(outerInner, bag)
	log.Println(len(outerBags))
	log.Println(innerCount)
}
