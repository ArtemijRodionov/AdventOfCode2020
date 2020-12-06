package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Exist struct{}

const ExpectedSum = 2020

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, val)
	}

	firstResult := make(map[int]Exist)
	secondResult := make(map[int]Exist)
	for i, first := range numbers {
		for j, second := range numbers {
			if i == j {
				continue
			}
			if first+second == ExpectedSum {
				firstResult[first*second] = Exist{}
			}

			for y, third := range numbers {
				if j == y || i == y {
					continue
				}
				if first+second+third == ExpectedSum {
					secondResult[first*second*third] = Exist{}
				}
			}
		}
	}
	for number := range firstResult {
		log.Println("First", number)
	}
	for number := range secondResult {
		log.Println("Second", number)
	}
}
