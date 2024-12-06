package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_puzzle(fpath string) (map[int][]int, [][]int) {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	odr := make(map[int][]int)
	lst := [][]int{}
	_ = lst

	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			continue
		} else if len(txt) == 5 {
			split := strings.Split(txt, "|")
			x, erra := strconv.Atoi(split[0])
			y, errb := strconv.Atoi(split[1])
			if erra != nil || errb != nil {
				log.Fatalf("Invalid number in line: %s", txt)
			}

			if _, exists := odr[x]; !exists {
				odr[x] = []int{}
			}
			odr[x] = append(odr[x], y)
		} else {
			split := strings.Split(txt, ",")
			tlst := []int{}
			for idx := range split {
				x, _ := strconv.Atoi(split[idx])
				tlst = append(tlst, x)
			}
			lst = append(lst, tlst)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return odr, lst
}

func solve_puzzle(odr map[int][]int, lst [][]int) int {

	total := 0
	for sli := range lst {
		mid := lst[sli][len(lst[sli])/2]
		valid := true
		for eli := range lst[sli] {
			if !valid {
				break
			}
			k := lst[sli][eli]
			tk, exists := odr[k]
			if !exists {
				continue
			}
			for tki := range tk {
				if !valid {
					break
				}
				for bi := eli - 1; bi >= 0; bi-- {
					if !valid {
						break
					}
					miv := tk[tki]
					inn := lst[sli][bi]
					if inn == miv {
						// Debug
						//fmt.Printf("INVALID INDEX: %d\n", sli)
						//fmt.Printf("Invalid : %d because %d|%d. See %v\n",
						//	lst[sli][eli], k, miv, lst[sli])
						valid = false
						break
					}
				}
			}
		}
		if valid {
			total += mid
		}
	}
	return total
}

func main() {

	odr, lst := read_puzzle("./small.txt")
	s := solve_puzzle(odr, lst)
	fmt.Printf("Part I (small): %d\n", s)
	if s != 143 {
		log.Fatal("Wrong Answer (part I (small))\n")
	}

	// 5681 too high
	odr2, lst2 := read_puzzle("./large.txt")
	s2 := solve_puzzle(odr2, lst2)

	fmt.Printf("Part I (large): %d\n", s2)
	if s2 != 5275 {
		log.Fatal("Wrong Answer (part I (small))\n")
	}

}
