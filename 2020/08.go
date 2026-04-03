package main

import "bufio"
import "os"
import "strconv"
import "fmt"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var is = make([]inst, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		n, _ := strconv.Atoi(line[4:])
		is = append(is, inst{ line[0:3], n })
	}
	var res [2]int
	res[0], _ = run(is)
	for i := 0; i < len(is); i++ {
		switch is[i].t {
		case "nop": is[i].t = "jmp"
		case "jmp": is[i].t = "nop"
		}
		var acc, ip = run(is)
		if ip == len(is) {
			res[1] = acc
			break
		}
		switch is[i].t {
		case "nop": is[i].t = "jmp"
		case "jmp": is[i].t = "nop"
		}
	}
	fmt.Println(res)
}

type inst struct { t string; n int }

func run(is []inst) (int, int) {
	var acc, ip = 0, 0
	var executed = make(map[int]bool)
	executed[len(is)] = true
	for {
		if executed[ip] { break }
		executed[ip] = true
		switch is[ip].t {
		case "acc": acc += is[ip].n; ip++
		case "nop": ip++
		case "jmp": ip += is[ip].n
		default: panic("wrong instruction " + is[ip].t + " at " + strconv.Itoa(ip))
		}
	}
	return acc, ip
}