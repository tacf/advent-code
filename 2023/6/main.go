package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"regexp"
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


func part2(file string) {
	log.Printf("####### Part 2: %s #######", file)
	lines, _ := readFileLines(file)
	space := regexp.MustCompile(`\s+`)
	time := space.ReplaceAllString(strings.Split(lines[0], ":")[1], "")
	distance := space.ReplaceAllString(strings.Split(lines[1], ":")[1], "")
	wins := 0
	maxTime, _ := strconv.Atoi(time)
	dist, _ := strconv.Atoi(distance)
	for j := 1; j <= maxTime; j++ {
		remainingTime := maxTime - j
		distTravelled := j * remainingTime
		if distTravelled > dist {
			//log.Printf("Remaining Time: %d, Distance Travelled: %d, Speed: %d", remainingTime, distTravelled, j)
			wins += 1
		}
	}
	log.Printf("Wins: %d", wins)
}

func part1(file string) {
	log.Printf("####### Part 1: %s #######", file)
	lines, _ := readFileLines(file)
	space := regexp.MustCompile(`\s+`)
	times := strings.Split(space.ReplaceAllString(lines[0], " "), " ")[1:]
	distance := strings.Split(space.ReplaceAllString(lines[1], " "), " ")[1:]
	wins := 0
	sum := 1
	for i, t := range times {
		log.Println("Race: ", t)
		maxTime, _ := strconv.Atoi(t)
		dist, _ := strconv.Atoi(distance[i])
		for j := 1; j <= maxTime; j++ {
			remainingTime := maxTime - j
			distTravelled := j * remainingTime
			if distTravelled > dist {
				wins += 1
			}
		}
		sum *= wins
		wins = 0
	}
	log.Printf("Result: %d", sum)
}

func run(file string) {
	part1(file)
	part2(file)
}

func main() {
	run("races.txt")
	os.Exit(0)
}
