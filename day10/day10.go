package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Trailhead struct {
	x int
	y int
}

type XY struct {
	x int
	y int
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func read_puzzle(fpath string) [][]int {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	topo := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		row := []int{}
		for c := range txt {
			i, _ := strconv.Atoi(string(txt[c]))
			row = append(row, i)
		}
		topo = append(topo, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return topo
}

func find_trailheads(puzzle [][]int) []Trailhead {

	trailheads := []Trailhead{}
	for y := range puzzle {
		for x := range puzzle[y] {
			if puzzle[y][x] == 0 {
				th := Trailhead{x: x, y: y}
				trailheads = append(trailheads, th)
			}
		}
	}
	return trailheads
}

func bound_check(puzzle [][]int, y int, x int) bool {
	ylen := len(puzzle)
	xlen := len(puzzle[0])
	xok := x >= 0 && x < xlen
	yok := y >= 0 && y < ylen
	return xok && yok
}

func dfs(topo [][]int, y int, x int, visited map[XY]bool, score int) int {

	if visited == nil {
		visited = make(map[XY]bool)
	}
	start := XY{x: x, y: y}
	visited[start] = true
	height := topo[y][x]
	//fmt.Printf("Visit: %d, (%d, %d)\n", height, x, y)
	for i := 0; i < 4; i++ {
		ny := directions[i][0] + y
		nx := directions[i][1] + x
		if !bound_check(topo, ny, nx) {
			continue
		}
		if topo[ny][nx] != height+1 {
			continue
		}
		neighbor := XY{x: nx, y: ny}
		_, ok := visited[neighbor]
		if !ok {
			score += dfs(topo, ny, nx, visited, 0)
		}
	}
	if height == 9 {
		score += 1
	}
	return score
}

func dfs_rating(
	topo [][]int,
	y int,
	x int,
	rating int,
	this_path []XY,
) int {

	node := XY{x: x, y: y}
	this_path = append(this_path, node)
	height := topo[y][x]
	if height == 9 {
		rating++
	} else {

		for i := 0; i < 4; i++ {
			ny := directions[i][0] + y
			nx := directions[i][1] + x
			if !bound_check(topo, ny, nx) {
				continue
			}
			neighbor := XY{x: nx, y: ny}
			for j := range this_path {
				if this_path[j] == neighbor {
					continue
				}
			}
			if topo[ny][nx] != height+1 {
				continue
			}
			//fmt.Printf("Visit: %d, (%d, %d)\n", height, x, y)
			rating = dfs_rating(topo, ny, nx, rating, this_path)
		}
	}
	return rating
}

func find_all_paths(
	topo [][]int,
	y int,
	x int,
) int {
	this_path := make([]XY, 0)
	return dfs_rating(topo, y, x, 0, this_path)
}

func solve_puzzle(filepath string, get_rating bool) int {
	total_score := 0
	topo := read_puzzle(filepath)
	trailheads := find_trailheads(topo)

	if !get_rating {
		for i := range trailheads {
			score := dfs(topo, trailheads[i].y, trailheads[i].x, nil, 0)
			//fmt.Printf("SCORE: %d", score)
			total_score += score
		}
		return total_score
	} else {
		total_rating := 0
		for i := range trailheads {
			rating := find_all_paths(topo, trailheads[i].y, trailheads[i].x)
			total_rating += rating
		}
		return total_rating
	}
}

func test_puzzle(filepath string, get_rating bool, expected int, description string) {

	solution := solve_puzzle(filepath, get_rating)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./small.txt", false, 36, "Part I (small):")
	test_puzzle("./large.txt", false, 611, "Part I (large):")
	test_puzzle("./tiny2.txt", true, 3, "Part I (tiny2):")
	test_puzzle("./small.txt", true, 81, "Part II (small):")
	test_puzzle("./large.txt", true, 1380, "Part II (large):")
}
