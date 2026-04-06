package main

import "bufio"
import "os"
import "fmt"

func main() {
    var res [2]int
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var expr = scanner.Text()
        var n, _ = eval(expr, 0)
        res[0] += n
        n, _ = evalMul(expr, 0)
        res[1] += n
    }
    fmt.Println(res[0], res[1])
}

func eval(expr string, i int) (int, int) {
    var val = 0
    var op = "+"
out:
    for ;i < len(expr); i++ {
        switch expr[i] {
        case '+': op = "+"
        case '*': op = "*"
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            var n = int(expr[i] - '0')
            switch op {
            case "+": val += n
            case "*": val *= n
            }
        case '(':
            var n = 0
            n, i = eval(expr, i + 1)
            switch op {
            case "+": val += n
            case "*": val *= n
            }
        case ')': break out
        }
    }
    return val, i
}

func evalSum(expr string, i int) (int, int) {
    var val = 0
out:
    for i < len(expr) {
        switch expr[i] {
        case '+': i++
        case '*': break out
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            val += int(expr[i] - '0')
            i++
        case '(':
            var n = 0
            n, i = evalMul(expr, i + 1)
            val += n
        case ')':
            break out
        case ' ':
            i++
        }
    }
    return val, i
}

func evalMul(expr string, i int) (int, int) {
    var val = 1
    val, i = evalSum(expr, i)
out:
    for i < len(expr) {
        switch expr[i] {
        case '*':
            var n = 0
            n, i = evalSum(expr, i + 1)
            val *= n
        case ')':
            i++
            break out
        case ' ': i++
        }
    }
    return val, i
}