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
	harmonics bool,
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
				if harmonics {
					antinodes[anta.y][anta.x]++
				}
				xoa, yoa := antenna_offset(anta, antb)
				out_of_bounds := false
				nxa := anta.x
				nya := anta.y
				for !out_of_bounds {
					nxa = nxa + xoa
					nya = nya + yoa
					out_of_bounds = !bound_check(puzzle, nya, nxa)
					if !out_of_bounds {
						antinodes[nya][nxa]++
					}
					if !harmonics {
						break
					}
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

func solve_puzzle(puzzle [][]string, harmonics bool) int {
	amap := antenna_map(puzzle)
	antinodes := make_antinodes_map(puzzle)
	compute_antinodes(puzzle, &antinodes, amap, harmonics)
	sol := uniq_antinode_pos(antinodes)
	return sol
}

func test_puzzle(filepath string, harmonics bool, expected int, description string) {
	puzzle := read_puzzle(filepath)
	solution := solve_puzzle(puzzle, harmonics)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./small.txt", false, 14, "Part I (small):")
	test_puzzle("./small_t.txt", true, 9, "Part I (small_t):")
	test_puzzle("./large.txt", false, 409, "Part II (large):")
	test_puzzle("./small.txt", true, 34, "Part II (small):")
	test_puzzle("./large.txt", true, 1308, "Part II (large):")
}
