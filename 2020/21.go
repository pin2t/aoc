package main

import "bufio"
import "os"
import "regexp"
import "strings"
import "fmt"
import "slices"

func main() {
    type food struct { ing, all []string }
    var scanner = bufio.NewScanner(os.Stdin)
    var reWord = regexp.MustCompile("[a-z]+")
    var alin = make(map[string]map[string]bool)
    var counts = make(map[string]int)
    for scanner.Scan() {
        var line = scanner.Text()
        var in, al = make(map[string]bool), make([]string, 0)
        var s = strings.Split(line, "contains")
        for _, word := range reWord.FindAllString(s[0], -1) {
            in[word] = true
            counts[word]++
        }
        for _, word := range reWord.FindAllString(s[1], -1) {
            al = append(al, word)
        }
        for _, a := range al {
            var nm = make(map[string]bool)
            if ii, found := alin[a]; found {
                for i, _ := range ii {
                    if in[i] { nm[i] = true }
                }
            } else {
                for k, v := range in { nm[k] = v }
            }
            alin[a] = nm
        }
    }
    var sum = 0
    for i, c := range counts {
        var found = false
        for _, ii := range alin {
            found = found || ii[i]
        }
        if !found { sum += c }
    }
    fmt.Print(sum)
    var reduced = true
    for reduced {
        reduced = false
        for a, i := range alin {
            if len(i) != 1 { continue }
            for v, _ := range i {
                for aa, ii := range alin {
                    if ii[v] && aa != a {
                        delete(ii, v)
                        reduced = true
                    }
                }
            }
        }
    }
    var sorted, res2 = make([]string, 0), ""
    for a, _ := range alin { sorted = append(sorted, a) }
    slices.Sort(sorted)
    for _, a := range sorted {
        if len(res2) > 0 { res2 += "," }
        for aa, _ := range alin[a] {
            res2 = res2 + aa
        }
    }
    fmt.Println("", res2)
}