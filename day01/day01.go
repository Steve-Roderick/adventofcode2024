package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func this_parse(fpath string) ([]int, []int) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ida := []int{}
	idb := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		numa, erra := strconv.Atoi(split[0])
		numb, errb := strconv.Atoi(split[1])
		if erra != nil || errb != nil {
			log.Fatalf("Invalid number in line: %s", scanner.Text())
		}
		ida = append(ida, numa)
		idb = append(idb, numb)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ida, idb
}

// part I
func total_distance_sort(fpath string) int {

	ida, idb := this_parse(fpath)
	sort.Ints(ida)
	sort.Ints(idb)
	var d int = 0
	for index := range ida {
		a := ida[index]
		b := idb[index]
		if a >= b {
			d += a - b
		} else {
			d += b - a
		}
	}
	return d
}

// part II
func sim_score(fpath string) int {
	ida, idb := this_parse(fpath)

	acc := 0
	freq := make(map[int]int)
	for _, element := range idb {
		freq[element] += 1
	}
	for _, element := range ida {
		acc += (element * freq[element])
	}
	return acc
}

func main() {
	small := total_distance_sort("./small.txt")
	fmt.Printf("Part I (small): %d\n", small)
	if small != 11 {
		log.Fatal("Wrong Answer")
	}

	large := total_distance_sort("./large.txt")
	fmt.Printf("Part I (large): %d\n", large)
	if large != 2113135 {
		log.Fatal("Wrong Answer")
	}

	small2 := sim_score("./small.txt")
	fmt.Printf("Part II (small): %d\n", small2)
	if small2 != 31 {
		log.Fatal("Wrong Answer")
	}

	large2 := sim_score("./large.txt")
	fmt.Printf("Part II (small): %d\n", large2)
	if large2 != 19097157 {
		log.Fatal("Wrong Answer")
	}
}
