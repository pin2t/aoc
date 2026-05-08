package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"

func main() {
    var reLiteral = regexp.MustCompile("^(\\d+):\\s\"(.*?)\"$")
    var reSeq = regexp.MustCompile("^(\\d+):\\s([\\d\\s]+)$")
    var reOr = regexp.MustCompile("^(\\d+):\\s([\\d\\s]+)|([\\d\\s]+)$")
    var scanner = bufio.NewScanner(os.Stdin)
    var matched = [2]int{ 0, 0 }
    for scanner.Scan() {
        var l = scanner.Text()
        if len(l) == 0 { break }
        var m [][]string
        if m = reLiteral.FindAllStringSubmatch(l, -1); len(m) != 0 {
            var n, _ = strconv.Atoi(m[0][1])
            rules[n] = rule{ m[0][2], make([][]int, 0) }
        } else if m = reSeq.FindAllStringSubmatch(l, -1); len(m) != 0 {
            var n, _ = strconv.Atoi(m[0][1])
            rules[n] = rule{ "", [][]int{ parse(m[0][2]) } }
        } else if m = reOr.FindAllStringSubmatch(l, -1); len(m) != 0 {
            var n, _ = strconv.Atoi(m[0][1])
            rules[n] = rule{ "", [][]int{ parse(m[0][2]), parse(m[1][0]) } }
        }
    }
    var messages = make([]string, 0)
    for scanner.Scan() {
        var msg = scanner.Text()
        messages = append(messages, msg)
        if m, i := matches(msg, 0, 0); m && i == len(msg) { matched[0]++ }
    }
    rules[8] = rule{ "", [][]int{ { 42 }, { 42, 8} } }
    rules[11] = rule{ "", [][]int{ { 42, 31 }, { 42, 11, 31 } } }
    for _, msg := range messages {
        var res = matches2(msg, 0, 0)
        for _, i := range res {
            if i == len(msg) { matched[1]++ }
        }
    }
    fmt.Println(matched)
}

var rules = make(map[int]rule)

type rule struct {
    val string
    rules [][]int
}

func matches(s string, ri int, si int) (bool, int) {
    var r = rules[ri]
    if r.val != "" {
        if si < len(s) && s[si:si+len(r.val)] == r.val {
            return true, si + len(r.val)
        } else {
            return false, 0
        }
    }
    for _, rs := range r.rules {
        var idx = si
        for _, ri := range rs {
            var m, i = false, 0
            if m, i = matches(s, ri, idx); !m {
                idx = -1
                break
            }
            idx = i
        }
        if idx != -1 {
            return true, idx
        }
    }
    return false, 0
}

func matches2(s string, ri int, si int) []int {
    var r = rules[ri]
    var results = make([]int, 0)
    if r.val != "" {
        if si < len(s) && s[si:si+len(r.val)] == r.val {
            results = append(results, si + len(r.val))
        }
    } else {
        for _, rs := range r.rules {
            results = apply(results, s, rs, 0, si)
        }
    }
    return results
}

func apply(results []int, s string, rs []int, ri int, si int) []int {
    if ri == len(rs) {
        return append(results, si)
    }
    var res = matches2(s, rs[ri], si)
    for _, idx := range res {
        results = apply(results, s, rs, ri + 1, idx)
    }
    return results
}

var reNums = regexp.MustCompile("\\d+")

func parse(s string) (res []int) {
    res = make([]int, 0)
    for _, s := range reNums.FindAllString(s, -1) {
        var n, _ = strconv.Atoi(s)
        res = append(res, n)
    }
    return
}
