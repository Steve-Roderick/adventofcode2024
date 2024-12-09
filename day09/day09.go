package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Block struct {
	id        int
	continous int
	moved     bool
}

func read_puzzle(fpath string) []int {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	disk_map := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		for c := range txt {
			i, _ := strconv.Atoi(string(txt[c]))
			disk_map = append(disk_map, i)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return disk_map
}

func blockify(disk_map []int) []Block {
	blocks := []Block{}
	idc := 0
	for i := range disk_map {
		if disk_map[i] == 0 {
			continue
		}
		block := Block{}
		block.continous = disk_map[i]
		block.moved = false
		is_file := (i % 2) == 0
		if is_file {
			block.id = idc
			idc += 1
		} else {
			block.id = -1
		}
		for j := 0; j < block.continous; j++ {
			next_block := block
			next_block.continous -= j
			blocks = append(blocks, next_block)
		}
	}
	return blocks
}

func index_last_file(blocks []Block) int {

	index := -1
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].id == -1 {
			continue
		}
		if blocks[i].moved {
			continue
		}
		index = i
		break
	}
	return index
}

func index_last_file_head(blocks []Block) int {
	index := index_last_file(blocks)
	if index == -1 {
		return index
	}
	first := -1
	for i := index; i >= 0; i-- {
		if blocks[i].id != blocks[index].id {
			first = i + 1
			break
		}
	}
	return first

}

func relocate_compress(pblocks *[]Block) bool {

	blocks := *pblocks
	start := index_last_file(blocks)
	//need := blocks[start].continous
	for i := range blocks {
		if blocks[i].id != -1 {
			continue
		}
		if blocks[i].continous >= 1 {
			blocks[i] = blocks[start]
			free_block := Block{}
			free_block.id = -1
			free_block.continous = 1
			blocks[start] = free_block

			return true
		}
	}
	return false
}

func relocate(pblocks *[]Block) bool {

	blocks := *pblocks
	start := index_last_file_head(blocks)

	if start == 0 || start == -1 {
		return false
	}
	need := blocks[start].continous
	for i := range blocks {
		if blocks[i].id != -1 {
			continue
		}
		if i > start {
			break
		}
		if blocks[i].continous >= need {
			free_block := Block{}
			free_block.id = -1
			free_block.continous = need
			for j := 0; j < need; j++ {
				new_block := blocks[start+j]
				new_block.moved = true
				free_block.continous = need - j
				blocks[i+j] = new_block
				blocks[start+j] = free_block
			}
			return true
		}
	}

	for i := 0; i < blocks[start].continous; i++ {
		blocks[start+i].moved = true
	}
	return true
}

func is_contiguous(blocks []Block) bool {
	found_first_empty := false
	for i := range blocks {
		if blocks[i].id == -1 {
			found_first_empty = true
		}
		if blocks[i].id != -1 && found_first_empty {
			return false
		}
	}
	return true
}

func checksum(blocks []Block) int {
	acc := 0
	for i := range blocks {

		/* 		if blocks[i].id == -1 {
		   			fmt.Printf(".")
		   		} else {
		   			fmt.Printf("%d", blocks[i].id)
		   		} */
		if blocks[i].id == -1 {
			continue
		}
		acc += i * blocks[i].id
	}
	//fmt.Printf("\n")
	return acc
}

func defrag(blocks []Block, compress bool) int {

	running := true
	if compress {
		for running && !is_contiguous(blocks) {
			running = relocate_compress(&blocks)
		}
	} else {
		for running {
			a := checksum(blocks)
			_ = a
			running = relocate(&blocks)
		}
	}
	cs := checksum(blocks)
	return cs
}

func solve_puzzle(filepath string, compress bool) int {
	disk_map := read_puzzle(filepath)
	disk_block := blockify(disk_map)
	sol := defrag(disk_block, compress)
	return sol
}

func test_puzzle(filepath string, compress bool, expected int, description string) {

	solution := solve_puzzle(filepath, compress)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./small.txt", true, 1928, "Part I (small):")
	test_puzzle("./large.txt", true, 6259790630969, "Part I (large):")
	test_puzzle("./small.txt", false, 2858, "Part II (small):")
	test_puzzle("./large.txt", false, 6289564433984, "Part II (large):")
}
