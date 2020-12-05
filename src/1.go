package main

import (
    "os"
    "bufio"
)

func Fn() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        println(scanner.Text())
    }
}

