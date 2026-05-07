package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var input string
    fmt.Scanln(&input)
    var s = input
    for i := 0; i < 100; i++ {
        var pick = [3]byte{ s[1], s[2], s[3] }
        var n = s[0]
        for {
            n--
            if n == '0' { n = '9' }
            if n != pick[0] && n != pick[1] && n != pick[2] { break }
        }
        s = s[0:1] + s[4:]
        var idx = strings.Index(s, string(n))
        s = s[0:idx + 1] + string(pick[0]) + string(pick[1]) + string(pick[2]) + s[idx+1:]
        s = s[1:] + string(s[0])
    }
    var idx = strings.Index(s, "1")
    fmt.Print(s[idx + 1:] + s[0:idx])
    type node struct { val int; next *node }
    var nums = make(map[int]*node)
    var val, _ = strconv.Atoi(string(input[0]))
    var cur = &node{ val, nil }
    var loop = cur
    nums[cur.val] = cur
    for i := 1; i < len(input); i++ {
        val, _ = strconv.Atoi(string(input[i]))
        var n = &node{ val, nil }
        loop.next = n
        loop = n
        nums[n.val] = n
    }
    for i := 10; i <= 1000000; i++ {
        var n = &node{ i, nil }
        loop.next = n
        loop = n
        nums[n.val] = n
    }
    loop.next = cur
    for i := 0; i < 10000000; i++ {
        loop = cur
        var pick  [3]int
        loop = loop.next
        pick[0] = loop.val
        loop = loop.next
        pick[1] = loop.val
        loop = loop.next
        pick[2] = loop.val
        var n = cur.next
        cur.next = loop.next
        var target = cur.val - 1
        if target == 0 { target = 1000000 }
        for pick[0] == target || pick[1] == target || pick[2] == target {
            target--
            if target == 0 { target = 1000000 }
        }
        loop = nums[target]
        var nx = loop.next
        loop.next = n
        loop = loop.next
        loop = loop.next
        loop = loop.next
        loop.next = nx
        cur = cur.next
    }
    fmt.Println("", nums[1].next.val * nums[1].next.next.val)
}
