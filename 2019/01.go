package main

import "fmt"

func main() {
    var sum = [2]int{0, 0}
    for {
        var mass = 0
        if _, err := fmt.Scanln(&mass); err != nil { break }
        sum[0] += mass/3 - 2
        for {
            mass = mass/3 - 2
            if mass <= 0 { break }
            sum[1] += mass
        }
    }
    fmt.Println(sum)
}
