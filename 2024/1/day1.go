package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func assertEqual[T comparable](a, b T) {
	if a != b {
		log.Fatalf("Assert Equals failed: %v != %v", a, b)
	}
}

func part1(c1, c2 []int) int {

	sort.Slice(c1, func(i, j int) bool { return c1[i] < c1[j] })
	sort.Slice(c2, func(i, j int) bool { return c2[i] < c2[j] })

	result := 0
	for i := range c1 {
		result = result + int(math.Abs(float64(c1[i]-c2[i])))
	}

	return result
}

func part2(c1, c2 map[int]int) int {

	result := 0
	for k, o1 := range c1 {
		if o2, ok := c2[k]; ok {
			result = result + (k * o1 * o2)
		}
	}

	return result
}

func parseInput1(filepath string) ([]int, []int) {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	c1 := []int{}
	c2 := []int{}

	scanner := bufio.NewScanner(file)
	rgx := regexp.MustCompile(`^(\d*)\s*(\d*)$`)
	for scanner.Scan() {
		line := scanner.Text()
		re := rgx.FindStringSubmatch(line)

		i1, _ := strconv.Atoi(re[1])
		i2, _ := strconv.Atoi(re[2])
		c1 = append(c1, i1)
		c2 = append(c2, i2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}

	return c1, c2
}

// func parseLines[T interface{}](filepath string, regex string, lineHandler func([]string) T) T {
//
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		log.Fatalf("Could not open file: %v", err)
// 	}
// 	defer file.Close()
//
// 	c1 := []int{}
// 	c2 := []int{}
//
// 	scanner := bufio.NewScanner(file)
// 	rgx := regexp.MustCompile(`^(\d*)\s*(\d*)$`)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		re := rgx.FindStringSubmatch(line)
//
// 		i1, _ := strconv.Atoi(re[1])
// 		i2, _ := strconv.Atoi(re[2])
// 		c1 = append(c1, i1)
// 		c2 = append(c2, i2)
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error occurred during scanning: %v", err)
// 	}
//
// 	return c1, c2
// }

func parseInput2(filepath string) (map[int]int, map[int]int) {
	m1 := map[int]int{}
	m2 := map[int]int{}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`^(\d*)\s*(\d*)$`)
	for scanner.Scan() {
		line := scanner.Text()
		re := regex.FindStringSubmatch(line)
		i1, _ := strconv.Atoi(re[1])
		i2, _ := strconv.Atoi(re[2])

		if val, ok := m1[i1]; ok {
			m1[i1] = val + 1
		} else {
			m1[i1] = 1
		}
		if val, ok := m2[i2]; ok {
			m2[i2] = val + 1
		} else {
			m2[i2] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}
	return m1, m2
}

func main() {
	x, y := parseInput1("input_example.txt")
	assertEqual(part1(x, y), 11)
	z, w := parseInput2("input_example.txt")
	assertEqual(part2(z, w), 31)
	fmt.Println("Example checks succeded")
	a, b := parseInput1("input.txt")
	fmt.Printf("Part1 result: %d\n", part1(a, b))
	c, d := parseInput2("input.txt")
	fmt.Printf("Part2 result: %d\n", part2(c, d))
}
