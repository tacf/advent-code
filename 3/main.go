package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func readFileLines(filePath string) ([]string, error) {

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func getLineNumbers(line string) [][]int {
	result := [][]int{}
	numberFoundState := false
	startIndex := -1

	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			numberFoundState = true
			if startIndex == -1 {
				startIndex = i
			}
		} else {
			if numberFoundState {
				numberFoundState = false
				result = append(result, []int{startIndex, i - 1})
				startIndex = -1
			}
		}
	}
	if numberFoundState{
		result = append(result, []int{startIndex, len(line) - 1})
	}
	return result
}

func checkChar(s string) bool {
	if strings.Contains("0123456789.", s){
		return false
	} else {
		return true
	}
}

func processLinePart1(lineb string, line string, linea string) int {
	result := 0

	lineNumbers := getLineNumbers(line)
	for _, numberIndexRange := range lineNumbers {
		charFound := false
		for i := numberIndexRange[0]; i <= numberIndexRange[1]; i++ {
			// first line position only has adjacent numberst to the right
			if i >= 0 && i < len(line)-1 {
				if checkChar(string(lineb[i+1])) || checkChar(string(linea[i+1])) || checkChar(string(line[i+1])) {
					charFound = true
				}
			}

			// Any position as corresponding one in other lines
			if checkChar(string(lineb[i])) || checkChar(string(linea[i])) {
				charFound = true
			}

			// last line position only has adjacent numberst to the left
			if i > 0 && i <= len(line) { 
				//println(string(lineb[i-1]), string(linea[i-1]), string(line[i-1]))
				if checkChar(string(lineb[i-1])) || checkChar(string(linea[i-1])) || checkChar(string(line[i-1])) {
					charFound = true
				}
			}


		}
		log.Printf("Number(%t): %s\n", charFound, line[numberIndexRange[0]:numberIndexRange[1]+1])
		if charFound {
			r, _ := strconv.Atoi(string(line[numberIndexRange[0]:numberIndexRange[1]+1]))
			result += r
			charFound = false
		} 
	}
	println("####")
	println(lineb)
	println(line)
	println(linea)
	println("####")
	return result
}


func part1(file string) {
	log.Printf("####### Part 1: %s #######", file)
	lines, _ := readFileLines(file)
	emptyLine := strings.Repeat(".", len(lines[0]))
	result := 0

	// Append empty lines to allow processing each line surrounded by two othe lines
	// Makes easier to apply same function to first and last line
	lines = append([]string{emptyLine}, lines...)
	lines = append(lines, emptyLine)

	for i := 1; i < len(lines)-1; i++ {
		result += processLinePart1(lines[i-1], lines[i], lines[i+1])
	}

	log.Printf("Result: %d", result)
}

func run(file string) {
	part1(file)
}

func main() {
	run("engine_parts.txt")
	os.Exit(0)
}
