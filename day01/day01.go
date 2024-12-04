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

func total_distance_sort(fpath string) int {
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
		numa, _ := strconv.Atoi(split[0])
		numb, _ := strconv.Atoi(split[1])
		ida = append(ida, numa)
		idb = append(idb, numb)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	slicea := ida[:]
	sliceb := idb[:]

	sort.Ints(slicea)
	sort.Ints(sliceb)
	var d int = 0
	for index := range slicea {
		a := slicea[index]
		b := sliceb[index]
		var c int = 0
		if a >= b {
			c = a - b
		} else {
			c = b - a
		}
		d += c
	}
	return d
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
}
