package main

import "bufio"
import "os"
import "strings"
import "regexp"
import "strconv"
import "fmt"

func main() {
    var mem, mem2 = make(map[int64]int64), make(map[int64]int64)
    var reMem = regexp.MustCompile(`^mem\[(\d+)\]\s=\s(\d+)$`)
    var ormask, andmask = int64(0), int64(-1)
    var mask string
    var fbits = 0
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var s = scanner.Text()
        if strings.HasPrefix(s, "mask = ") {
            mask = s[7:]
            fbits = 0
            ormask, andmask = int64(0), int64(-1)
            var bit = int64(1) << (len(s) - 8)
            for _, c := range s[7:] {
                switch c {
                case '0': andmask ^= bit
                case '1': ormask |= bit
                case 'X': fbits++
                }
                bit >>= 1
            }
        } else {
            var parts = reMem.FindStringSubmatch(s)
            var addr, _ =  strconv.ParseInt(parts[1], 10, 64)
            var val, _ =  strconv.ParseInt(parts[2], 10, 64)
            mem[addr] = (val | ormask) & andmask
            var f = 1 << fbits - 1
            for f >= 0 {
                var fb = f
                var faddr = addr | ormask
                for i := len(mask) - 1; i >= 0; i-- {
                    if mask[i] == 'X' {
                        faddr &= ^(1 << (len(mask) - 1 - i))
                        faddr |= int64(fb & 1) << (len(mask) - 1 - i)
                        fb >>= 1
                    }
                }
                mem2[faddr] = val
                f--
            }
        }
    }
    fmt.Println(sum(mem), sum(mem2))
}

func sum(m map[int64]int64) (res int64) {
    for _, val := range m {
        res += val
    }
    return
}