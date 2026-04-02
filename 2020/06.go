package main

import "bufio"
import "os"
import "fmt"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var letters = make(map[rune]bool)
	var all = make(map[rune]bool)
	var res = []int{ 0, 0 }
	for scanner.Scan() {
		var l = scanner.Text()
		if len(l) == 0 {
			res[0] += len(letters)
			res[1] += len(all)
			letters = make(map[rune]bool)
			all = make(map[rune]bool)
		} else {
			var p = make(map[rune]bool)
			if len(letters) == 0 {
				for _, c := range l { all[c] = true }
			}
			for _, c := range l {
				letters[c] = true
				p[c] = true
			}
			for k := range all {
				if !p[k] {
					delete(all, k)
				}
			}
		}
	}
	res[0] += len(letters)
	res[1] += len(all)
	fmt.Println(res)
}
