package main

import (
    "os"
    "bufio"
    "strconv"
    "log"
)

type Exist struct {}
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

    result := make(map[int]Exist)
    for i, first := range numbers {
        for j, second := range numbers {
            if i == j { continue }
            if first + second == ExpectedSum {
                result[first * second] = Exist{}
            }
        }
    }
    for number := range result {
        log.Println(number)
    }
}

