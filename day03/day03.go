package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type descent struct {
	// tokens
	m  bool
	u  bool
	l  bool
	op bool
	n0 bool
	se bool
	n1 bool
	cp bool
	// states
	n0c int
	n1c int
	n0s string
	n1s string
}

func null_descent(ts *descent) {
	ts.m = false
	ts.u = false
	ts.l = false
	ts.op = false
	ts.n0 = false
	ts.se = false
	ts.n1 = false
	ts.cp = false

	ts.n0c = 0
	ts.n1c = 0
	ts.n0s = ""
	ts.n1s = ""

}

// Token parser with some state to accept variable number of numbers (1-3).
func parse(memory string) int {

	ts := descent{}

	acc := 0
	reset := true
	for index := range memory {
		if reset {
			null_descent(&ts)
			reset = false
		}

		c := rune(memory[index])
		d := unicode.IsDigit(rune(c))
		cc := string(c) // debug.
		_ = cc

		if c == 'm' && !ts.m {
			ts.m = true
			continue
		}
		if c == 'u' && !ts.u && ts.m {
			ts.u = true
			continue
		}
		if c == 'l' && !ts.l && ts.u && ts.m {
			ts.l = true
			continue
		}
		if c == '(' && !ts.op && ts.l && ts.u && ts.m {
			ts.op = true
			continue
		}
		if !ts.se && d && ts.op && ts.l && ts.u && ts.m {
			ts.n0c += 1
			if ts.n0c > 3 {
				reset = true
				continue
			}
			ts.n0s = ts.n0s + string(c)
			ts.n0 = true
			continue
		}
		if c == ',' && !ts.se && ts.n0 && ts.op && ts.l && ts.u && ts.m {
			ts.se = true
			continue
		}
		if !ts.cp && d && ts.se && ts.n0 && ts.op && ts.l && ts.u && ts.m {
			ts.n1c += 1
			if ts.n1c > 3 {
				reset = true
				continue
			}
			ts.n1s = ts.n1s + string(c)
			ts.n1 = true
			continue
		}
		if c == ')' && !ts.cp && ts.n1 && ts.se && ts.n0 && ts.op && ts.l && ts.u && ts.m {
			ts.cp = true
		}
		if ts.cp && ts.n1 && ts.se && ts.n0 && ts.op && ts.l && ts.u && ts.m {
			numa, _ := strconv.Atoi(ts.n0s)
			numb, _ := strconv.Atoi(ts.n1s)
			acc += numa * numb
			// dbg.
			//fmt.Printf("mul(%d,%d)\n", numa, numb)
			reset = true
		} else {
			reset = true
		}
	}

	return acc
}

func read_memory(fpath string) string {

	data, err := os.ReadFile(fpath)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal("Oups")
	}
	longString := string(data)
	return longString
}

func main() {
	small := read_memory("./small.txt")
	small_1 := parse(small)
	fmt.Printf("Part I (small): %d\n", small_1)
	if small_1 != 161 {
		log.Fatal("Wrong Answer")
	}

	large := read_memory("./large.txt")
	large_1 := parse(large)
	fmt.Printf("Part I (large): %d\n", large_1)
	if large_1 != 166630675 {
		log.Fatal("Wrong Answer")
	}
}
