package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var namedDigits = map[string]int{
	"zero": 0,
	"one":   1,
	"two": 2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":  6,
	"seven":  7,
	"eight":  8,
	"nine":  9,
}

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

func getFirstNumber(line string) int {
	for i, char := range line {
		if char >= '0' && char <= '9' {
			return i
		}
	}
	return -1
}

func getLastNumber(line string) int {
	return len(line) - (getFirstNumber(reverse(line))) - 1
}


func getFirstNamedDigit(line string) (int, string) {
	foundNamedDigits := make(map[string]int)
	for key, _ := range namedDigits {
		index := strings.Index(line, key)
		if index != -1 {
			foundNamedDigits[key] = index
		}
	}
	minIndex := -1
	var name string
	for n, index := range foundNamedDigits {
		if minIndex == -1 || index < minIndex {
			minIndex = index
			name = n
		}
	}
	return minIndex, name
}

func getLastNamedDigit(line string) (int, string) {
	foundNamedDigits := make(map[string]int)
	for key, _ := range namedDigits {
		index := strings.LastIndex(line, key)
		if index != -1 {
			foundNamedDigits[key] = index
		}
	}
	maxIndex := -1
	var name string
	for n, index := range foundNamedDigits {
		if maxIndex == -1 || index > maxIndex {
			maxIndex = index
			name = n
		}
	}
	return maxIndex, name
}

func reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}

func run(file string) {

	log.Printf("####### Calculating file: %s #######", file)	
	lines, _ := readFileLines(file)
	result := 0
	for _, line := range lines {
		firstNumberIndex := getFirstNumber(line)
		lastNumberIndex := getLastNumber(line)
		firstNamedNumberIndex, firstNamedNumberName := getFirstNamedDigit(line)
		lastNamedNumberIndex, lastNamedNumberName := getLastNamedDigit(line)

		var firstNumber int
		var lastNumber int
		if firstNumberIndex == -1 || (firstNamedNumberIndex < firstNumberIndex && firstNamedNumberIndex != -1)  {
			firstNumber = namedDigits[firstNamedNumberName]
		} else {
			firstNumber = int(line[firstNumberIndex]-'0')
		}

		if lastNumberIndex == len(line) || (lastNamedNumberIndex > lastNumberIndex && lastNamedNumberIndex != -1) {
			lastNumber = namedDigits[lastNamedNumberName]
		} else {
			lastNumber = int(line[lastNumberIndex]-'0')
		}

		result += firstNumber*10 + lastNumber

	}
	log.Printf("Result: %d", result)
}

func main() {
	run("calibration_doc.txt")
	os.Exit(0)
}