package main

import "fmt"
import "regexp"
import "strconv"
import "strings"

func main() {
    var ts, ids = 0, ""
    fmt.Scanln(&ts)
    fmt.Scanln(&ids)
    var reNum = regexp.MustCompile("\\d+")
    var m, mid = 1000000, 0
    for _, sid := range reNum.FindAllString(ids, -1) {
        var id, _ = strconv.Atoi(sid)
        if id - ts % id < m {
            m = id - ts % id
            mid = id
        }
    }
    fmt.Print(m * mid)
    var start, mul = int64(0), int64(0)
    for i, sid := range strings.Split(ids, ",") {
        if sid == "x" { continue }
        var id, _ = strconv.ParseInt(sid, 10, 64)
        if start == 0 {
            start = id
            mul = id
            continue
        }
        var b = id
        var goal = start + int64(i)
        for b != goal && goal % id > 0 {
            b = (goal / id + 1) * id
            for b > goal { goal += mul }
        }
        var goal2 = goal + mul
        for b != goal2 && goal2 % id > 0{
            b = (goal2 / id + 1) * id
            for b > goal2 { goal2 += mul}
        }
        start = goal - int64(i)
        mul = goal2 - goal
    }
    fmt.Println("", start)
}
