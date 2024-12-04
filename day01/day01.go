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

func main() {

	file, err := os.Open("./large.txt")
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
	fmt.Printf("Part I: %d\n", d)
}
