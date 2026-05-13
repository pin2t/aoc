package main

import "bufio"
import "os"
import "strings"
import "fmt"

func main() {
    var scanner = bufio.NewScanner(os.Stdin)
    var orbit = make(map[string]string)
    for scanner.Scan() {
        var link = strings.Split(scanner.Text(), ")")
        orbit[link[1]] = link[0]
    }
    var total = 0
    for _, link := range orbit {
        var c = 1
        for link != "COM" {
            link = orbit[link]
            c++
        }
        total += c
    }
    fmt.Print(total)
    var path = make(map[string]int)
    var o, c = orbit["YOU"], 0
    for o != "COM" {
        path[o] = c
        o = orbit[o]
        c++
    }
    o, c = "SAN", 0
    for o != "COM" {
        o = orbit[o]
        if v, found := path[o]; found {
            fmt.Println("", v + c)
            break
        }
        c++
    }
}
