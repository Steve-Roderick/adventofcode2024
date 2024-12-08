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

func solve_puzzle(puzzle [][]string, sx int, sy int, cycled *[][][]int) (int, int, bool) {

	var movement = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	mi := 0
	x := sx
	y := sy

	steps := 0
	running := true
	in_cycle := false
	for running && !in_cycle {
		if cycled == nil {
			puzzle[x][y] = "X"
		}
		next_x := x + movement[mi][0]
		next_y := y + movement[mi][1]
		if !bound_check(puzzle, next_x, next_y) {
			running = false
			in_cycle = false
			break
		}
		if puzzle[next_x][next_y] == "#" {
			mi = move_mi(mi)
			continue
		}

		if cycled != nil {
			for i := range (*cycled)[x][y] {
				if (*cycled)[x][y][i] == mi {
					in_cycle = true
					break
				}
			}
			if !in_cycle {
				for i := range (*cycled)[x][y] {
					if (*cycled)[x][y][i] == -1 {
						(*cycled)[x][y][i] = mi
					}
				}
			}
		}
		x = next_x
		y = next_y
		steps++
	}

	acc := 0
	for rowi := range puzzle {
		for coli := range puzzle[rowi] {
			if puzzle[rowi][coli] == "X" {
				acc++
			}
		}
	}
	return acc, steps, in_cycle
}

// 10 steps
// step {1,10}
// skip step at x, y == start
// There is a problem when we call this function which I kludged over.
// Technically we should know exactly how many times to call this func
// based on how many obstacles (steps) are in the puzzle.
func find_path_obstacle(puzzle [][]string, x int, y int, steps int, step int) (int, int) {

	_ = steps
	acc := 0
	for i := range puzzle {
		for j := range puzzle[i] {
			if i == x && j == y {
				continue
			}
			if puzzle[i][j] == "X" {
				acc++
				if acc == step {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

func find_cycles(puzzle [][]string, x int, y int, steps int) int {

	acc := 0
	tx := -1
	ty := -1
	for stepi := range steps {
		if stepi != 0 {
			puzzle[tx][ty] = "X"
		}
		cycled := init_cycle_detect(puzzle, x, y)
		tx, ty = find_path_obstacle(puzzle, x, y, steps, stepi+1)
		if tx == -1 {
			break
		}
		puzzle[tx][ty] = "#"
		_, _, in_cycle := solve_puzzle(puzzle, x, y, &cycled)
		if in_cycle {
			acc++
			//fmt.Printf("Cycle with: %d, %d\n", tx, ty)
		}
	}
	return acc
}

func init_cycle_detect(puzzle [][]string, x int, y int) [][][]int {

	cycled := [][][]int{}
	_ = cycled
	rows := len(puzzle)
	cols := len(puzzle[0])

	cycled = make([][][]int, rows)
	for i := range cycled {
		cycled[i] = make([][]int, cols)
		for j := range cycled[i] {
			cycled[i][j] = make([]int, 4)
			for k := range cycled[i][j] {
				cycled[i][j][k] = -1
			}
		}
	}
	_ = x
	_ = y
	return cycled
}

func main() {
	// Small Problem
	puzzle1, x, y := read_puzzle("./small.txt")
	s, steps, _ := solve_puzzle(puzzle1, x, y, nil)
	fmt.Printf("Part I (small): %d\n", s)
	if s != 41 {
		log.Fatal("Wrong")
	}

	t := find_cycles(puzzle1, x, y, steps)
	fmt.Printf("Part II (small): %d\n", t)
	if t != 6 {
		log.Fatal("Wrong")
	}

	// Large Problem.
	puzzle2, x2, y2 := read_puzzle("./large.txt")
	s2, steps2, _ := solve_puzzle(puzzle2, x2, y2, nil)
	fmt.Printf("Part II (large): %d\n", s2)
	if s2 != 5242 {
		log.Fatal("Wrong")
	}

	p2 := find_cycles(puzzle2, x2, y2, steps2)
	fmt.Printf("Part II (small): %d\n", p2)
	if p2 != 1424 {
		log.Fatal("Wrong")
	}
}
