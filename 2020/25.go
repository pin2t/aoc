package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    var val, loop = 1, 1
    for loop = 1; val != b; loop++ { val = (val * 7) % 20201227 }
    val, loop = 1, loop - 1
    for l := 0; l < loop; l++ { val = (val * a) % 20201227 }
    fmt.Println(val)
}
