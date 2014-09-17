package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    inputFile, err := os.Open("2014_fall_1.sch")
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := inputFile.Close(); err != nil {
            panic(err)
        }
    }()
 
    scanner := bufio.NewScanner(inputFile)

    var i int = 0

    for scanner.Scan() {
        if i % (4 * 4 * 7 * 6 + 1) == 0 {
            fmt.Println(scanner.Text())
        }
        i++
    }
}