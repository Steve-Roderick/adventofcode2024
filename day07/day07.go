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
		calibration := Calibration{}
		calibration.test_value = test_value
		calibration.constants = constants
		calibration.solved = -1
		//fmt.Println(test_value)
		//fmt.Println(constants)
		//fmt.Println(calibration)
		puzzle = append(puzzle, calibration)

	}

	//fmt.Println(puzzle)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return puzzle
}

func solve_puzzle(puzzle []Calibration) int {

	for i := range puzzle {
		calibration := puzzle[i]
		solved := 0
		n_equations := int(math.Pow(2, float64(len(calibration.constants)-1)))
		for j := 0; j < n_equations; j++ {
			eval_value := calibration.constants[0]
			for k := range calibration.constants[1:] {
				v := calibration.constants[k+1]
				// Use LSB bit to signal operation.
				// 0b01 ==> *
				// 0b00 ==> +
				// For the kth operation, use the opt code for this equation.
				// Shift back k bits to get significant bit.
				optcode := ((0b01 << k) & j) >> k
				//fmt.Println(optcode)
				if optcode == 0 {
					eval_value = eval_value * v
				} else if optcode == 1 {
					eval_value = eval_value + v
				} else {
					log.Fatal("BIT SHIFT MAGIC DOES NOT COMPUTE")
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
	sol := solve_puzzle(puzzle)
	fmt.Printf("Small (part I): %d\n", sol)
	if sol != 3749 {
		log.Fatal("Wrong")
	}

	puzzle2 := read_puzzle("large.txt")
	sol2 := solve_puzzle(puzzle2)
	fmt.Printf("Large (part I): %d\n", sol2)
	if sol2 != 2664460013123 {
		log.Fatal("Wrong")
	}

	fmt.Printf("Ok\n")

}
