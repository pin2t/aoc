package main

import "bufio"
import "os"
import "regexp"
import "strings"
import "fmt"
import "strconv"

var nest = make(map[string]map[string]int)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var reBags = regexp.MustCompile(`(\d+ )?(\w+\s\w+) bags?`)
	for scanner.Scan() {
		var line = scanner.Text()
		var matches = reBags.FindAllStringSubmatch(line, -1)
		var b = matches[0][2]
		for i := 1; i < len(matches); i++ {
			if matches[i][0] == "no other bags" {
				continue
			}
			if nest[b] == nil {
				nest[b] = make(map[string]int)
			}
			j, _ := strconv.Atoi(strings.TrimSpace(matches[i][1]))
			nest[b][matches[i][2]] = j
		}
	}
	var contains = make(map[string]bool)
	var total = 0
	{
		var q = make([]string, 0)
		q = append(q, "shiny gold")
		for len(q) > 0 {
			var bag = q[0]
			q = q[1:]
			if contains[bag] {
				continue
			}
			contains[bag] = true
			for b, nb := range nest {
				for n := range nb {
					if n == bag {
						q = append(q, b)
					}
				}
			}
		}
	}
	{
		type state struct {
			bag string
			mul int
		}
		var q = make([]state, 0)
		q = append(q, state{"shiny gold", 1})
		for len(q) > 0 {
			var st = q[0]
			q = q[1:]
			for name, cnt := range nest[st.bag] {
				total += cnt * st.mul
				q = append(q, state{name, cnt * st.mul})
			}
		}
	}
	fmt.Println(len(contains)-1, total)
}
