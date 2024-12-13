package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Plot struct {
	garden string
	gid    int
	x      int
	y      int
	fc     int
}

func init_plot(y int, x int, garden string) Plot {
	p := Plot{}
	p.garden = garden
	p.x = x
	p.y = y
	p.gid = -1
	p.fc = -1
	return p
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func bound_check(plots [][]Plot, y int, x int) bool {
	ylen := len(plots)
	xlen := len(plots[0])
	xok := x >= 0 && x < xlen
	yok := y >= 0 && y < ylen
	return xok && yok
}

func read_puzzle(fpath string) [][]Plot {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	plots := [][]Plot{}
	scanner := bufio.NewScanner(file)
	y := -1
	for scanner.Scan() {
		y++
		txt := scanner.Text()
		row := make([]Plot, 0)
		for x := range txt {
			p := init_plot(y, x, string(txt[x]))
			row = append(row, p)
		}
		plots = append(plots, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return plots
}

func greasy_fence_estimate(
	pmap map[int][]*Plot,
	fmap map[int]int,
) int {
	acc := 0
	for i := range pmap {
		area := len(pmap[i])
		perm := fmap[i]
		acc += (area * perm)
	}
	return acc
}

func fence_count(plots [][]Plot, y int, x int) int {
	fences := 0
	for i := 0; i < 4; i++ {
		ny := directions[i][0] + y
		nx := directions[i][1] + x
		if !bound_check(plots, ny, nx) {
			fences++
			continue
		}
		if plots[ny][nx].garden != plots[y][x].garden {
			fences++
			continue
		}
	}
	return fences
}

func BFS(pplots *[][]Plot, start_y int, start_x int, gid int) []*Plot {
	group := make([]*Plot, 0)
	plots := *pplots
	visited := make(map[*Plot]bool)

	start := &(plots[start_y][start_x])
	start.gid = gid
	this_garden := start.garden
	group = append(group, start)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*Plot)
		queue.Remove(element)

		for i := 0; i < 4; i++ {
			ny := directions[i][0] + node.y
			nx := directions[i][1] + node.x
			if !bound_check(plots, ny, nx) {
				continue
			}
			neighbor := &(plots[ny][nx])
			c0 := this_garden == neighbor.garden
			c1 := !visited[neighbor]

			if c0 && c1 {
				neighbor.gid = gid
				group = append(group, neighbor)
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}
	return group
}

func solve_puzzle(filepath string) int {
	plots := read_puzzle(filepath)

	pmap := map[int][]*Plot{}
	fmap := make(map[int]int, 0)
	gid := -1
	for y := range plots {
		for x := range plots[y] {
			if plots[y][x].gid == -1 {
				gid++
				group := BFS(&plots, y, x, gid)
				pmap[gid] = group

				acc := 0
				for i := range group {
					group[i].gid = gid
					c := fence_count(plots, group[i].y, group[i].x)
					group[i].fc = c
					acc += c
				}
				fmap[gid] = acc
			}
		}
	}
	cost := greasy_fence_estimate(pmap, fmap)
	return cost
}

func test_puzzle(filepath string, expected int, description string) {

	solution := solve_puzzle(filepath)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./tiny.txt", 140, "Part I (tiny):")
	test_puzzle("./small.txt", 1930, "Part I (small):")
	test_puzzle("./large.txt", 1387004, "Part I (large):")
}
