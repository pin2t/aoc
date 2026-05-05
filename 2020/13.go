package main

import "fmt"
import "regexp"
import "strconv"

func main() {
    var ts, ids = 0, ""
    fmt.Scanln(&ts)
    fmt.Scanln(&ids)
    var reNum = regexp.MustCompile("\\d+")
    var m, mid = 1000000, 0
    for _, id := range reNum.FindAllString(ids, -1) {
        var nid, _ = strconv.Atoi(id)
        if nid - ts % nid < m {
            m = nid - ts % nid
            mid = nid
        }
    }
    fmt.Println(m * mid)
}
