package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	Part1()
	Part2()

}

func Part1() {
	races := ParseInput(puzzleInput)
	errorMargin := 1
	for _, race := range races {
		winningCount := CalcNumberOfRoutes(race.Time, race.Distance)
		errorMargin *= winningCount
	}

	fmt.Println("Part 1: ", errorMargin)
}

func Part2() {
	races := ParseInputPart2(puzzleInput)
	start := time.Now()
	t, d := races[0].Time, races[0].Distance
	winningCount := CalcNumberOfRoutes(t, d)

	fmt.Println("Part 2: ", winningCount)
	fmt.Println("took: ", time.Now().Sub(start))
}

func runRace(waitTime, totalTime int) int {
	return waitTime * (totalTime - waitTime)
}

func CalcNumberOfRoutes(maxTime, winningDistance int) int {
	wins := 0
	for waitTime := 0; waitTime <= winningDistance; waitTime++ {
		distance := runRace(waitTime, maxTime)
		if distance > winningDistance {
			wins++
		}
	}

	return wins
}

func ParseInput(input string) []Race {
	numberRegex := regexp.MustCompile(`(\d+)`)

	matches := numberRegex.FindAllStringSubmatch(input, -1)

	halfway := len(matches) / 2
	races := []Race{}
	for i := 0; i < halfway; i++ {
		race := Race{
			Time:     mustAtoi(matches[i][0]),
			Distance: mustAtoi(matches[halfway+i][0]),
		}

		races = append(races, race)

	}

	return races
}

func ParseInputPart2(input string) []Race {
	numberRegex := regexp.MustCompile(`(\d+)`)

	matches := numberRegex.FindAllStringSubmatch(input, -1)

	halfway := len(matches) / 2
	var time, distance string
	races := []Race{}
	for i := 0; i < halfway; i++ {

		time = fmt.Sprintf("%s%s", time, matches[i][0])
		distance = fmt.Sprintf("%s%s", distance, matches[halfway+i][0])

	}

	races = append(races, Race{
		Time:     mustAtoi(time),
		Distance: mustAtoi(distance),
	})

	return races
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

const testInput = `Time:      7  15   30
Distance:  9  40  200`

const puzzleInput = `Time:        56     97     77     93
Distance:   499   2210   1097   1440`
