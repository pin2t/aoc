package main

import "bufio"
import "os"
import "fmt"

func main() {
    var sz, i, area = 5, 1, 0
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        for _, c := range scanner.Text() {
            if c == '#' { area |= i }
            i *= 2
        }
    }
    var areas = make(map[int]bool)
    areas[area] = true
    for {
        var next int
        for i := 0; i < sz * sz; i++ {
            var nb = 0
            if i % sz > 0 && (area & (1 << (i - 1)) > 0) { nb++ }
            if i % sz < (sz - 1) && (area & (1 << (i + 1)) > 0) { nb++ }
            if i / sz > 0 && (area & (1 << (i - sz)) > 0) { nb++ }
            if i / sz < (sz - 1) && (area & (1 << (i + sz)) > 0) { nb++ }
            if (area & (1 << i) > 0) && nb == 1 { next |= 1 << i }
            if (area & (1 << i) == 0) && (nb == 1 || nb == 2) { next |= 1 << i }
        }
        if areas[next] {
            fmt.Print(next)
            break
        }
        areas[next] = true
        area = next
    }
}
