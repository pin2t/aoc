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
    var matched = 0
    for scanner.Scan() {
        var l = scanner.Text()
        if len(l) == 0 { break }
        var m = reLiteral.FindAllStringSubmatch(l, -1)
        if len(m) != 0 {
            var n, _ = strconv.Atoi(m[0][1])
            rules[n] = ruleLiteral{ m[0][2] }
        } else {
            m = reSeq.FindAllStringSubmatch(l, -1)
            if len(m) != 0 {
                var n, _ = strconv.Atoi(m[0][1])
                rules[n] = ruleSeq{ parse(m[0][2]) }
            } else {
                m = reOr.FindAllStringSubmatch(l, -1)
                if len(m) != 0 {
                    var n, _ = strconv.Atoi(m[0][1])
                    rules[n] = ruleOr{ parse(m[0][2]), parse(m[1][0]) }
                }
            }
        }
    }
    var messages = make([]string, 0)
    for scanner.Scan() {
        var msg = scanner.Text()
        messages = append(messages, msg)
        if m, left := rules[0].match(msg); m && len(left) == 0 {
            matched++
        }
    }
    fmt.Println(matched)
}

var rules = make(map[int]rule)

type rule interface{
    match(s string) (bool, string)
}

type ruleLiteral struct { lit string }

func (r ruleLiteral) match(s string) (bool, string) {
    if s[0:len(r.lit)] == r.lit { return true, s[len(r.lit):] }
    return false, s
}

type ruleSeq struct { seq []int }

func (r ruleSeq) match(s string) (bool, string) {
    var m = true
    var left = s
    for _, it := range r.seq {
        m, left = rules[it].match(left)
        if !m { return false, left }
    }
    return true, left
}

type ruleOr struct { seq1 []int; seq2 []int }

func (r ruleOr) match(s string) (bool, string) {
    var rl = ruleSeq{ r.seq1 }
    var m, left = rl.match(s)
    if m { return m, left }
    rl = ruleSeq{ r.seq2 }
    m, left = rl.match(s)
    if m { return m, left }
    return false, ""
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
