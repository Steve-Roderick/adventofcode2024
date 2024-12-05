package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var XMAS = "XMAS"
var SAMX = "SAMX"

func read_puzzle(fpath string) []string {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := -1
	puzzle := []string{}
	for scanner.Scan() {
		row += 1
		puzzle = append(puzzle, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return puzzle
}

// easy cases
func scan_xmas_samx(str string) int {
	acc := 0
	acc += strings.Count(str, XMAS)
	acc += strings.Count(str, SAMX)
	return acc
}

func puzzle_solver(puzzle []string) int {

	// Handle each case in seperate loop
	// for the god of debug.

	// Horizontal
	acc := 0
	for _, element := range puzzle {
		acc += scan_xmas_samx(element)
	}

	// Vertical
	for j := range puzzle[0] {
		var sb strings.Builder
		for i := range puzzle {
			sb.WriteString(string(puzzle[i][j]))
		}
		acc += scan_xmas_samx(sb.String())

	}
	rows := len(puzzle)
	cols := len(puzzle[0])

	diag_i := []string{}
	// Diag I
	for i := range puzzle {
		r := i
		c := 0
		var sb strings.Builder
		for r < rows && c < cols {
			sb.WriteString(string(puzzle[r][c]))
			r += 1
			c += 1
		}
		diag_i = append(diag_i, sb.String())
	}

	// Diag II
	//for i := range puzzle[0][1:] {
	for i := 1; i < cols; i++ {
		r := 0
		c := i
		var sb strings.Builder
		for r < rows && c < cols {
			sb.WriteString(string(puzzle[r][c]))
			r += 1
			c += 1
		}
		diag_i = append(diag_i, sb.String())
	}

	// Diag III
	for i := range puzzle {
		r := i
		c := cols - 1
		var sb strings.Builder
		for r < rows && c >= 0 {
			sb.WriteString(string(puzzle[r][c]))
			r += 1
			c -= 1
		}
		diag_i = append(diag_i, sb.String())
	}

	// Diag IV
	for i := len(puzzle[0]) - 2; i > -1; i-- {
		r := 0
		c := i
		var sb strings.Builder
		for r < rows && c >= 0 {
			sb.WriteString(string(puzzle[r][c]))
			r += 1
			c -= 1
		}
		diag_i = append(diag_i, sb.String())
	}

	for _, element := range diag_i {
		acc += scan_xmas_samx(element)
	}
	return acc
}

func main() {

	small_puzzle := read_puzzle("./small.txt")
	c := puzzle_solver(small_puzzle)
	fmt.Printf("Part I (small): %d\n", c)
	if c != 18 {
		log.Fatal("Wrong Answer (part I (small))\n")
	}

	large_puzzle := read_puzzle("./large.txt")
	c2 := puzzle_solver(large_puzzle)
	fmt.Printf("Part I (large): %d\n", c2)
	if c2 != 2549 {
		log.Fatal("Wrong Answer)\n")
	}
}
