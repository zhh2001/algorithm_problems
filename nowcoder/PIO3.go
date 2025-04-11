package main

import (
    "fmt"
)

func main() {
    var a, b int
    for {
        n, _ := fmt.Scan(&a, &b)
        if n == 0 {
            break
        } else {
            fmt.Printf("%d\n", a + b)
        }
    }
}