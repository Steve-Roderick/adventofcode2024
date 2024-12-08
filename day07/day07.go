package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	test_value int
	constants  []int
	solved     int
}

func read_puzzle(fpath string) []Calibration {

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	puzzle := []Calibration{}
	scanner := bufio.NewScanner(file)
	max := 0
	for scanner.Scan() {
		txt := scanner.Text()

		split := strings.Split(txt, ":")
		test_value, _ := strconv.Atoi(split[0])
		split2 := strings.Split(split[1], " ")
		constants := []int{}
		for i := range split2[1:] {
			constant, _ := strconv.Atoi(split2[i+1])
			constants = append(constants, constant)
		}
		if len(constants) > max {
			max = len(constants)
		}
		calibration := Calibration{}
		calibration.test_value = test_value
		calibration.constants = constants
		calibration.solved = -1
		puzzle = append(puzzle, calibration)

	}
	fmt.Printf("MAX: %d\n", max)

	//fmt.Println(puzzle)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return puzzle
}

func solve_puzzle(puzzle []Calibration, num_operators int) int {

	for i := range puzzle {
		calibration := puzzle[i]
		solved := 0
		n_equations := int(math.Pow(float64(num_operators), float64(len(calibration.constants)-1)))
		for j := 0; j < n_equations; j++ {
			eval_value := calibration.constants[0]
			current_state := j
			for k := range calibration.constants[1:] {
				v := calibration.constants[k+1]
				// Use LSB bit to signal operation.
				// 0b01 ==> *
				// 0b00 ==> +
				// For the kth operation, use the opt code for this equation.
				// Shift back k bits to get significant bit.
				// Same same for three operator but use two bits.
				var optcode int
				if num_operators == 2 {
					optcode = ((0b01 << k) & j) >> k
				} else if num_operators == 3 {
					optcode = current_state % 3
					current_state /= 3
				}
				//fmt.Println(optcode)
				if optcode == 0 {
					eval_value = eval_value * v
				} else if optcode == 1 {
					eval_value = eval_value + v
				} else if optcode == 2 {
					eval_value_str := strconv.Itoa(eval_value)
					v_str := strconv.Itoa(v)
					eval_value, _ = strconv.Atoi(eval_value_str + v_str)
				} else {
					log.Fatal("BIT SHIFT MAGIC DOES NOT COMPUTE")
				}
				if eval_value > calibration.test_value {
					break
				}
			}
			if eval_value == calibration.test_value {
				solved++
			}
		}
		calibration.solved = solved
		puzzle[i] = calibration
	}
	acc := 0
	for i := range puzzle {
		if puzzle[i].solved > 0 {
			acc += puzzle[i].test_value
		}
	}
	return acc
}

func main() {
	// Small Problem
	puzzle := read_puzzle("small.txt")
	sol := solve_puzzle(puzzle, 2)
	fmt.Printf("Small (part I): %d\n", sol)
	if sol != 3749 {
		log.Fatal("Wrong")
	}

	// Large Problem
	puzzle2 := read_puzzle("large.txt")
	sol2 := solve_puzzle(puzzle2, 2)
	fmt.Printf("Large (part I): %d\n", sol2)
	if sol2 != 2664460013123 {
		log.Fatal("Wrong")
	}

	// Small Problem
	puzzle3 := read_puzzle("small.txt")
	sol3 := solve_puzzle(puzzle3, 3)
	fmt.Printf("Small (part II): %d\n", sol3)
	if sol3 != 11387 {
		log.Fatal("Wrong")
	}

	// Small Problem
	puzzle4 := read_puzzle("large.txt")
	sol4 := solve_puzzle(puzzle4, 3)
	fmt.Printf("Large (part II): %d\n", sol4)
	if sol4 != 426214131924213 {
		log.Fatal("Wrong")
	}
	fmt.Println("Ok")
}
