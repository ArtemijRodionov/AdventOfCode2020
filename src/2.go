package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Fn() {
	scanner := bufio.NewScanner(os.Stdin)
	validNumber := 0
	validNumber2 := 0
	for scanner.Scan() {
		policy := strings.Split(scanner.Text(), " ")
		if len(policy) > 3 {
			log.Fatal("Wrong input", policy)
		}

		policyRange := strings.Split(policy[0], "-")
		from, err := strconv.Atoi(policyRange[0])
		fail(err)
		to, err := strconv.Atoi(policyRange[1])
		fail(err)

		char := rune(policy[1][:1][0])
		password := policy[2]

		count := 0
		isValid := false
		for i, pchar := range password {
			if char == pchar {
				count++
			}
			idx := i + 1
			if (idx == from || idx == to) && pchar == char {
				isValid = !isValid
			}
		}

		if from <= count && to >= count {
			validNumber++
		}
		if isValid {
			validNumber2++
		}
	}
	log.Println(validNumber)
	log.Println(validNumber2)
}
