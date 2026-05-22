package main

import "regexp"
import "bufio"
import "os"
import "strconv"
import "strings"
import "fmt"
import "math"

func main() {
    type quantity struct { n int; mat string }
    type material struct { n int; from []quantity }
    var react = make(map[string]material)
    var re = regexp.MustCompile("(\\d+)\\s(\\w+)")
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var gr = re.FindAllStringSubmatch(scanner.Text(), -1)
        var qs = make([]quantity, 0)
        var toq = func (in []string) quantity {
            var n, _ = strconv.Atoi(in[1])
            return quantity{ n, strings.TrimSpace(in[2]) }
        }
        for i := 0; i < len(gr) - 1; i++ {
            qs = append(qs, toq(gr[i]))
        }
        var n, _ = strconv.Atoi(gr[len(gr) - 1][1])
        react[gr[len(gr) - 1][2]] = material{ n, qs }
    }
    var ore = func (fuel int64) (res int64) {
        var counts = map[string]int64{ "FUEL": fuel }
        for {
            var done = true
            var prod, req = make(map[string]int64), make(map[string]int64)
            for m, q := range counts {
                prod[m] = int64(react[m].n) * q
                for _, rq := range react[m].from {
                    req[rq.mat] += q * int64(rq.n)
                }
            }
            res = 0
            for m, q := range req {
                if m == "ORE" {
                    res += q
                } else if prod[m] < q {
                    done = false
                    counts[m] += int64(math.Ceil(float64(q - prod[m]) * 1.0 / float64(react[m].n)))
                }
            }
            if done { break }
        }
        return
    }
    fmt.Print(ore(1))
    var l, r = int64(1), int64(1000000000000)
    for l < r {
        var m = (l + r) / 2
        if ore(m) <= 1000000000000 { l = m + 1 } else { r = m - 1 }
    }
    fmt.Println("", l)
}
