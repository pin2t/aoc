package main

import "fmt"

func main() {
    var w, h = 25, 6
    var input string
    fmt.Scanln(&input)
    var layers = make([]string, 0)
    var min0 = "000000000000000000000000000000000000000000000000000000000000000000000000000000"
    for i := 0; i < len(input); i += w * h {
        var l = input[i:i+w*h]
        layers = append(layers, l)
        if cnt(l, '0') < cnt(min0, '0') { min0 = l }
    }
    fmt.Println(cnt(min0, '1') * cnt(min0, '2'))
    var pixels = map[byte]string{ '1': "*", '0': "." }
    for r := 0; r < h; r++ {
        for c := 0; c < w; c++ {
            var l = 0
            for ;l < len(layers) && layers[l][r * w + c] == '2'; l++ {}
            fmt.Print(pixels[layers[l][r * w + c]])
        }
        fmt.Println()
    }
}

func cnt(layer string, ch rune) int {
    var res = 0
    for _, c := range layer {
        if c == ch { res++ }
    }
    return res
}
