package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

func read_puzzle(fpath string) []int {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stones := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		split := strings.Split(txt, " ")
		for i := range split {
			n, _ := strconv.Atoi(string(split[i]))
			stones = append(stones, n)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stones
}

func the_trimmer(to_trim string) int {
	var n int
	if (len(to_trim)) == 1 {
		n, _ = strconv.Atoi(to_trim)
	} else {
		r := strings.TrimLeftFunc(to_trim, func(r rune) bool {
			return r == '0'
		})
		n, _ = strconv.Atoi((r))
	}
	return n
}

func the_splitter(alpha string) (int, int) {
	mid := len(alpha) / 2
	a1 := alpha[:mid]
	a2 := alpha[mid:]
	n1 := the_trimmer(a1)
	n2 := the_trimmer(a2)
	return n1, n2
}

func blink(stones []int) []int {

	stones2 := make([]int, 0)
	slices.Reverse(stones)
	for i := len(stones) - 1; i >= 0; i-- {
		n := stones[i]
		a := strconv.Itoa(stones[i])
		if n == 0 {
			stones2 = append(stones2, 1)

		} else if len(a)%2 == 0 {
			n1, n2 := the_splitter(a)
			stones2 = append(stones2, n1)
			stones2 = append(stones2, n2)
			continue
		} else {
			stones2 = append(stones2, n*2024)
		}
		stones = stones[:len(stones)-1]
		if i%1000 == 0 {
			runtime.GC()
		}
	}

	return stones2
}

// yikes!
func blink2(stone_smoke map[int]int) map[int]int {

	stones := make([]int, 0)
	for k := range stone_smoke {
		stones = append(stones, k)
	}
	stone_rock := make(map[int]int, 0)
	for k, v := range stone_smoke {
		stone_rock[k] = v
	}
	slices.Sort(stones)
	for k := range stones {
		n := stones[k]
		a := strconv.Itoa(n)
		count := stone_rock[n]
		if n == 0 {
			_, e := stone_smoke[1]
			if !e {
				stone_smoke[1] = stone_rock[0]
			} else {
				stone_smoke[1] += stone_rock[0]
			}

		} else if len(a)%2 == 0 {
			n1, n2 := the_splitter(a)

			_, e1 := stone_smoke[n1]
			if !e1 {
				stone_smoke[n1] = stone_rock[n]
			} else {
				stone_smoke[n1] += count
			}
			_, e2 := stone_smoke[n2]
			if !e2 {
				stone_smoke[n2] = stone_rock[n]
			} else {
				stone_smoke[n2] += count
			}

		} else {
			new := n * 2024
			_, e := stone_smoke[new]
			if !e {
				stone_smoke[new] = stone_rock[n]
			} else {
				stone_smoke[new] += stone_rock[n]
			}

		}
		stone_smoke[n] -= count
		if stone_smoke[n] <= 0 {
			delete(stone_smoke, n)
		}
	}
	return stone_smoke
}

func squash(stones []int) map[int]int {
	stone_smoke := make(map[int]int, 0)
	for i := range stones {
		_, ok := stone_smoke[stones[i]]
		if !ok {
			stone_smoke[stones[i]] = 1
		} else {
			stone_smoke[stones[i]] += 1
		}
	}
	return stone_smoke
}

func solve_puzzle(filepath string, blink_n int) int {
	stones := read_puzzle(filepath)

	for i := 0; i < blink_n; i++ {
		//fmt.Println(stones)
		stones = blink(stones)

		//fmt.Printf("%d, %d\n", i, len(stones))
	}
	//fmt.Println(stones)
	return len(stones)
}

func solve_puzzle2(filepath string, blink_n int) int {
	stones := read_puzzle(filepath)
	stone_smoke := squash(stones)
	for i := 0; i < blink_n; i++ {

		stone_smoke = blink2(stone_smoke)
		for k, v := range stone_smoke {
			if v == 0 {
				delete(stone_smoke, k)
			}
		}
		//fmt.Println(stone_smoke)
		//fmt.Printf("%d, %d\n", i, len(stones))
	}
	//fmt.Println(stones)

	acc := 0
	for k := range stone_smoke {
		acc += stone_smoke[k]
	}
	return acc
}

func test_puzzle(filepath string, blink_n int, expected int, description string) {

	solution := solve_puzzle(filepath, blink_n)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func test_puzzle2(filepath string, blink_n int, expected int, description string) {

	solution := solve_puzzle2(filepath, blink_n)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./tiny.txt", 6, 22, "Part I (tiny 6):")
	test_puzzle("./tiny.txt", 25, 55312, "Part I (small 25):")
	test_puzzle("./large.txt", 25, 239714, "Part I (large 25):")
	test_puzzle("./large.txt", 75, 239714, "Part I (large 25):")

	test_puzzle2("./tiny.txt", 1, 3, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 2, 4, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 3, 5, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 4, 9, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 5, 13, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 6, 22, "Part II (tiny 6):")
	test_puzzle2("./tiny.txt", 25, 55312, "Part II (small 25):")
	test_puzzle2("./large.txt", 25, 239714, "Part II (large 25):")
	test_puzzle2("./large.txt", 75, 284973560658514, "Part II (large 75):")
}
