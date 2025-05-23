package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Claw struct {
	// a button
	ax int
	ay int

	// b button
	bx int
	by int

	// prize
	px int
	py int

	// children
	lc *Claw
	rc *Claw

	cost  int
	mcost int
	found bool

	x int
	y int
}

var BUTTON_A_TXT = "Button A:"
var BUTTON_B_TXT = "Button B:"
var PRIZE_TXT = "Prize"

var LARGE_INT = 9223372036854775807

func extract_xy_button(line string) (int, int) {
	// Regular expression to match the two integers
	re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)

	// Find the matches
	matches := re.FindStringSubmatch(line)

	var x int
	var y int
	if len(matches) == 3 {
		// Convert the matched strings to integers
		x, _ = strconv.Atoi(matches[1])
		y, _ = strconv.Atoi(matches[2])
	}
	return x, y
}

func extract_xy_prize(line string) (int, int) {
	// Regular expression to match the two integers
	re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	// Find the matches
	matches := re.FindStringSubmatch(line)

	var x int
	var y int
	if len(matches) == 3 {
		// Convert the matched strings to integers
		x, _ = strconv.Atoi(matches[1])
		y, _ = strconv.Atoi(matches[2])
	}
	return x, y
}

func read_puzzle(fpath string) []Claw {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	claws := []Claw{}
	scanner := bufio.NewScanner(file)
	y := -1
	the_claw := Claw{}
	the_claw.lc = nil
	the_claw.rc = nil
	the_claw.x = 0
	the_claw.y = 0
	the_claw.mcost = LARGE_INT
	for scanner.Scan() {
		y++
		txt := scanner.Text()

		if txt == "" {
			claw := the_claw
			claws = append(claws, claw)
			continue
		}

		if txt[:len(BUTTON_A_TXT)] == BUTTON_A_TXT {
			x, y := extract_xy_button(txt)
			the_claw.ax = x
			the_claw.ay = y
			continue
		}

		if txt[:len(BUTTON_B_TXT)] == BUTTON_B_TXT {
			x, y := extract_xy_button(txt)
			the_claw.bx = x
			the_claw.by = y
			continue
		}

		if txt[:len(PRIZE_TXT)] == PRIZE_TXT {
			x, y := extract_xy_prize(txt)
			the_claw.px = x
			the_claw.py = y
			continue
		}

	}
	claw := the_claw
	claws = append(claws, claw)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return claws
}

func push_a(claw *Claw) {
	claw.x += claw.ax
	claw.y += claw.ay
	claw.cost += 3
}

func push_b(claw *Claw) {
	claw.x += claw.bx
	claw.y += claw.by
	claw.cost += 1
}

func rewind_a(claw *Claw) {
	claw.x -= claw.ax
	claw.y -= claw.ay
	claw.cost -= 3
}

func rewind_b(claw *Claw) {
	claw.x -= claw.bx
	claw.y -= claw.by
	claw.cost -= 1
}

func is_sat(claw *Claw) bool {
	if claw.x == claw.px && claw.y == claw.py {
		claw.found = true
		if claw.cost < claw.mcost {
			claw.mcost = claw.cost
		}
		return true
	} else {
		return false
	}
}

func is_bust(claw *Claw) bool {
	if claw.x > claw.px {
		return true
	}
	if claw.y > claw.py {
		return true
	}
	return false
}

func search_puzzle(claw *Claw, depth int) int {

	depth++
	fmt.Printf("DEPTH: %d\n", depth)
	if is_sat(claw) {
		return claw.cost
	}
	if is_bust(claw) {
		return -1
	}

	push_a(claw)
	cost_left := search_puzzle(claw, depth)

	if cost_left == -1 {
		rewind_a(claw)
	}

	push_b(claw)
	cost_right := search_puzzle(claw, depth)

	if cost_right == -1 {
		rewind_b(claw)
	}
	return -1
}

func solve_puzzle(filepath string) int {
	claws := read_puzzle(filepath)
	acc := 0
	for i := range claws {
		a := search_puzzle(&(claws[i]), 0)
		_ = a
		if claws[i].found {
			acc += claws[i].mcost
		}
	}

	return acc
}

func test_puzzle(filepath string, expected int, description string) {

	solution := solve_puzzle(filepath)
	fmt.Printf("%s %d\n", description, solution)
	if solution != expected {
		log.Fatalf("Wrong: expected %d, got %d\n", expected, solution)
	}
}

func main() {
	test_puzzle("./tiny.txt", 140, "Part I (small):")

}
