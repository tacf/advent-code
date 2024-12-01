package main

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func processLinePart1(line string) int {
	processCard := strings.Split(line, ":")
	cardValues := strings.Split(processCard[1], "|")
	winningNumbers := strings.Split(strings.Replace(strings.Trim(cardValues[0], " "), "  ", " ", -1), " ")
	myNumbers := strings.Split(strings.Replace(strings.Trim(cardValues[1], " "), "  ", " ", -1), " ")
	// log.Println(winningNumbers, " | ", myNumbers)
	cardResult := 0
	for _, n := range myNumbers {
		for _, w := range winningNumbers {
			if strings.Trim(w, "") == strings.Trim(n, "") {
				if cardResult == 0 {
					cardResult = 1
				} else {
					cardResult = cardResult * 2
				}
			}
		}
	}
	// log.Println(processCard[0], " | ", cardResult)
	return cardResult
}

func processLinePart2(line string) int {
	processCard := strings.Split(line, ":")
	cardValues := strings.Split(processCard[1], "|")
	winningNumbers := strings.Split(strings.Replace(strings.Trim(cardValues[0], " "), "  ", " ", -1), " ")
	myNumbers := strings.Split(strings.Replace(strings.Trim(cardValues[1], " "), "  ", " ", -1), " ")
	// log.Println(winningNumbers, " | ", myNumbers)
	cardResult := 0
	for _, n := range myNumbers {
		for _, w := range winningNumbers {
			if strings.Trim(w, "") == strings.Trim(n, "") {
				cardResult += 1
			}
		}
	}
	//log.Println(processCard[0], " | ", cardResult)
	return cardResult
}


func processLinesPart2(lines []string, start int, end int) int{
	result := 0
	for i := start; i < end ; i++ {
		lineResult := processLinePart2(lines[i])
		if lineResult > 0 {
			result += 1
			if i == len(lines) - 1 {
				return result
			} else {
				result += processLinesPart2(lines[i+1:], 0, lineResult)
			}
		} else {
			result += 1
		}
	}
	return result
}

func part2(file string) {
	log.Printf("####### Part 2: %s #######", file)
	lines, _ := readFileLines(file)
	result := processLinesPart2(lines, 0, len(lines))
	log.Printf("Result: %d", result)

}

func part1(file string) {
	log.Printf("####### Part 1: %s #######", file)
	lines, _ := readFileLines(file)
	result := 0
	for _, line := range lines {
		result += processLinePart1(line)
	}
	log.Printf("Result: %d", result)
}

func run(file string) {
	part1(file)
	part2(file)
}

func main() {
	run("cards.txt")
	os.Exit(0)
}
