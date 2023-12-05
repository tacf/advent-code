package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func processLinePart1(line string) (int, bool) {
	processGame := strings.Split(line, ":")
	gameNumber, _ := strconv.Atoi(strings.Split(processGame[0], " ")[1])
	gameRounds := strings.Split(processGame[1], ";")
	gameResult := true
	for _, round := range gameRounds {
		colors := strings.Split(round, ",")
		for _, colorString := range colors {
			colorDetails := strings.Split(colorString, " ")
			colorName := colorDetails[2]
			colorNumber, ok := strconv.Atoi(colorDetails[1])
			//log.Printf("Color: %d, %s", colorNumber, colorName)
			switch colorName {
			case "red":
				if ok != nil || colorNumber > 12 {
					gameResult = gameResult && false
				}
			case "green":
				if ok != nil || colorNumber > 13 {
					gameResult = gameResult && false
				}
			case "blue":
				if ok != nil || colorNumber > 14 {
					gameResult = gameResult && false
				}

			}
		}
	}
	return gameNumber, gameResult
}

func processLinePart2(line string) (int, map[string]int) {
	processGame := strings.Split(line, ":")
	gameNumber, _ := strconv.Atoi(strings.Split(processGame[0], " ")[1])
	gameRounds := strings.Split(processGame[1], ";")
	gameResult := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}
	for _, round := range gameRounds {
		colors := strings.Split(round, ",")
		for _, colorString := range colors {
			colorDetails := strings.Split(colorString, " ")
			colorName := colorDetails[2]
			colorNumber, _ := strconv.Atoi(colorDetails[1])
			//log.Printf("Color: %d, %s", colorNumber, colorName)
			value, _ := gameResult[colorName]
			if value < colorNumber {
				gameResult[colorName] = colorNumber
			}
		}
	}
	log.Printf("Game: %d, Result: %d, %d, %d", gameNumber, gameResult["red"], gameResult["green"], gameResult["blue"])
	return gameNumber, gameResult
}

func part2(file string) {
	log.Printf("####### Part 2: %s #######", file)
	lines, _ := readFileLines(file)
	result := 0
	for _, line := range lines {
		_, gameResult := processLinePart2(line)
		//log.Printf("Game: %d, Result: %d", gameNumber, (gameResult["red"] * gameResult["green"] * gameResult["blue"]))
		result = result + (gameResult["red"] * gameResult["green"] * gameResult["blue"])
	}
	log.Printf("Result: %d", result)
}

func part1(file string) {
	log.Printf("####### Part 1: %s #######", file)
	lines, _ := readFileLines(file)
	result := 0
	for _, line := range lines {
		gameNumber, gameResult := processLinePart1(line)
		//log.Printf("Game: %d, Result: %t", gameNumber, gameResult)
		if gameResult == true {
			result = result + gameNumber
		}
	}
	log.Printf("Result: %d", result)
}

func run(file string) {
	part1(file)
	part2(file)
}

func main() {
	run("games.txt")
	os.Exit(0)
}
