package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_report(fpath string) [][]int {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	on_level := -1
	for scanner.Scan() {
		on_level += 1

		if on_level >= len(reports) {
			reports = append(reports, []int{})
		}
		split := strings.Fields(scanner.Text())
		for _, element := range split {
			level, _ := strconv.Atoi(element)
			reports[on_level] = append(reports[on_level], level)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reports
}

func level_is_safe(levels []int) bool {
	direction := 0
	unsafe := false
	for idy := range levels {
		if idy == 0 {
			continue
		}
		before := levels[idy-1]
		current := levels[idy]
		if direction == 0 {
			if before > current {
				direction = -1
			} else if before < current {
				direction = 1
			}
		}
		diff := 0
		if direction == 1 {
			unsafe = current <= before
			diff = current - before
		} else if direction == -1 {
			unsafe = current >= before
			diff = before - current
		} else {
			unsafe = true
		}
		unsafe = unsafe || !(diff >= 1 && diff <= 3)
		if unsafe {
			break
		}
	}
	return !unsafe
}

func report_is_safe(reports [][]int, damper bool) int {
	safe := 0
	for idx := range reports {
		ok := level_is_safe(reports[idx])
		if ok {
			safe += 1
		} else if damper {
			for index := range reports[idx] {
				cpy := make([]int, 0)
				cpy = append(cpy, reports[idx][:index]...)
				cpy = append(cpy, reports[idx][index+1:]...)
				ok := level_is_safe(cpy)
				cpy = nil
				if ok {
					safe += 1
					break
				}
			}
		}
	}
	return safe
}

func main() {

	small_reports := parse_report("./small.txt")
	safe1 := report_is_safe(small_reports, false)
	fmt.Printf("Part I (small): %d\n", safe1)
	if safe1 != 2 {
		log.Fatal("Wrong Answer")
	}

	large_reports := parse_report("./large.txt")
	safe2 := report_is_safe(large_reports, false)
	fmt.Printf("Part I (large): %d\n", safe2)
	if safe2 != 526 {
		log.Fatal("Wrong Answer")
	}
	fmt.Printf("\n")

	safe3 := report_is_safe(small_reports, true)
	fmt.Printf("Part II (small): %d\n", safe3)
	if safe3 != 4 {
		log.Fatal("Wrong Answer")
	}
	safe4 := report_is_safe(large_reports, true)
	fmt.Printf("Part II (large): %d\n", safe4)

}
