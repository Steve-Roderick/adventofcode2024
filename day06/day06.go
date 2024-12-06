package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_puzzle(fpath string) ([][]string, int, int) {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzle := [][]string{}
	_ = puzzle

	x := -1
	y := -1
	rowi := -1
	for scanner.Scan() {
		rowi++
		txt := scanner.Text()
		row := []string{}
		for i := range txt {
			if txt[i] == '^' {
				x = rowi
				y = i
			}
			row = append(row, string(txt[i]))
		}
		puzzle = append(puzzle, row)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return puzzle, x, y
}

func bound_check(puzzle [][]string, x int, y int) bool {
	ylen := len(puzzle[0])
	xlen := len(puzzle)
	xok := x >= 0 && x < xlen
	yok := y >= 0 && y < ylen
	return xok && yok
}

func move_mi(mi int) int {
	mi++
	if mi > 3 {
		mi = 0
	}
	return mi
}

func solve_puzzle(puzzle [][]string, sx int, sy int) int {

	var movement = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	mi := 0
	x := sx
	y := sy

	running := true
	for running {
		puzzle[x][y] = "X"
		next_x := x + movement[mi][0]
		next_y := y + movement[mi][1]
		if !bound_check(puzzle, next_x, next_y) {
			running = false
			break
		}
		if puzzle[next_x][next_y] == "#" {
			mi = move_mi(mi)
			continue
		}

		x = next_x
		y = next_y
	}

	acc := 0
	for rowi := range puzzle {
		for coli := range puzzle[rowi] {
			if puzzle[rowi][coli] == "X" {
				acc++
			}
		}
	}
	return acc
}

func main() {
	puzzle1, x, y := read_puzzle("./small.txt")
	s := solve_puzzle(puzzle1, x, y)
	fmt.Printf("Part I (small): %d\n", s)
	if s != 41 {
		log.Fatal("Wrong")
	}

	puzzle2, x2, y2 := read_puzzle("./large.txt")
	s2 := solve_puzzle(puzzle2, x2, y2)
	fmt.Printf("Part I (large): %d\n", s2)
	if s2 != 5242 {
		log.Fatal("Wrong")
	}

}
