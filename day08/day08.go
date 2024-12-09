package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Antenna struct {
	freq string
	x    int
	y    int
}

func read_puzzle(fpath string) [][]string {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	puzzle := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []string{}
		txt := scanner.Text()
		for c := range txt {
			row = append(row, string(txt[c]))
		}
		puzzle = append(puzzle, row)
	}
	//fmt.Println(puzzle)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return puzzle
}

func antenna_map(puzzle [][]string) map[string][]Antenna {
	m := make(map[string][]Antenna)
	for y := range puzzle {
		for x := range puzzle[y] {
			c := puzzle[y][x]
			if c != "." && c != "#" {
				antenna := Antenna{}
				antenna.freq = c
				antenna.x = x
				antenna.y = y
				_, ok := m[c]
				if !ok {
					lst := []Antenna{}
					lst = append(lst, antenna)
					m[c] = lst
				} else {
					m[c] = append(m[c], antenna)
				}
			}
		}
	}
	return m
}

func make_antinodes_map(puzzle [][]string) [][]int {
	antinodes := [][]int{}
	for y := range puzzle {
		row := []int{}
		for x := range puzzle[y] {
			_ = x
			row = append(row, 0)
		}
		antinodes = append(antinodes, row)
	}
	return antinodes
}

func antenna_offset(a Antenna, b Antenna) (int, int) {
	x := a.x - b.x
	y := a.y - b.y
	return x, y
}

func bound_check(puzzle [][]string, y int, x int) bool {
	ylen := len(puzzle)
	xlen := len(puzzle[0])
	xok := x >= 0 && x < xlen
	yok := y >= 0 && y < ylen
	return xok && yok
}

func compute_antinodes(
	puzzle [][]string,
	antinodes_p *[][]int,
	antennas map[string][]Antenna,
) {
	antinodes := *antinodes_p
	// For each frequency
	for _, lst := range antennas {
		// For each combination of antennas
		for i := range lst {
			for j := range lst {
				if i == j {
					continue
				}
				anta := lst[i]
				antb := lst[j]
				xoa, yoa := antenna_offset(anta, antb)
				nxa := anta.x + xoa
				nya := anta.y + yoa
				if bound_check(puzzle, nya, nxa) {
					antinodes[nya][nxa]++
				}
			}
		}
	}
}

func uniq_antinode_pos(antinodes [][]int) int {
	acc := 0
	for i := range antinodes {
		for j := range antinodes {
			if antinodes[i][j] != 0 {
				acc++
			}
		}
	}
	return acc
}

func solve_puzzle(puzzle [][]string) int {
	amap := antenna_map(puzzle)
	antinodes := make_antinodes_map(puzzle)
	compute_antinodes(puzzle, &antinodes, amap)
	sol := uniq_antinode_pos(antinodes)
	return sol
}

func main() {
	// Small Problem
	puzzle1 := read_puzzle("./small.txt")
	sol1 := solve_puzzle(puzzle1)
	fmt.Printf("Part I (small) %d\n", sol1)
	if sol1 != 14 {
		log.Fatal("Wrong")
	}

	// Large Problem
	puzzle2 := read_puzzle("./large.txt")
	sol2 := solve_puzzle(puzzle2)
	fmt.Printf("Part I (large) %d\n", sol2)
	if sol2 != 409 {
		log.Fatal("Wrong")
	}
}
