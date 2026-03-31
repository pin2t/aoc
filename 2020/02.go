package main

import "bufio"
import "os"
import "regexp"
import "fmt"
import "strconv"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var re = regexp.MustCompile("(\\d+)-(\\d+)\\s(.*):\\s(.*)")
	var res, res2 = 0, 0
	for scanner.Scan() {
		var parts = re.FindAllStringSubmatch(scanner.Text(), -1)
		var l = rune(parts[0][3][0])
		from, _ := strconv.Atoi(parts[0][1])
		to, _ := strconv.Atoi(parts[0][2])
		var n = 0
		for _, c := range parts[0][4] {
			if c == l {
				n++
			}
		}
		if from <= n && n <= to {
			res++
		}
		var n2 = 0
		if rune(parts[0][4][from-1]) == l {
			n2++
		}
		if rune(parts[0][4][to-1]) == l {
			n2++
		}
		if n2 == 1 {
			res2++
		}
	}
	fmt.Println(res, res2)
}
