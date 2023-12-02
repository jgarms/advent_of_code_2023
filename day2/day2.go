package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result, err := partOne("sample_input.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Part one sample result:", result)

	result, err = partOne("input.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Part one result:", result)

	result, err = partTwo("sample_input.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Part two sample result:", result)

	result, err = partTwo("input.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Part two result:", result)
}

func partOne(inputFilePath string) (int, error) {
	games, err := parseInputPartOne(inputFilePath)
	if err != nil {
		return 0, err
	}
	var sum int
	for _, game := range games {
		if isGamePossible(game) {
			sum += game.ID
		}
	}
	return sum, nil
}

func partTwo(inputFilePath string) (int, error) {
	games, err := parseInputPartOne(inputFilePath)
	if err != nil {
		return 0, err
	}
	var sum int
	for _, game := range games {
		minCubeSet := computeMinimumCubeSet(game)
		sum += power(minCubeSet)
	}
	return sum, nil
}

func power(cubeSet CubeSet) int {
	return cubeSet.Red * cubeSet.Blue * cubeSet.Green
}

func computeMinimumCubeSet(game Game) CubeSet {
	minCubeSet := game.CubeSets[0]
	for i := 1; i < len(game.CubeSets); i++ {
		otherCubeSet := game.CubeSets[i]
		if otherCubeSet.Red > minCubeSet.Red {
			minCubeSet.Red = otherCubeSet.Red
		}
		if otherCubeSet.Blue > minCubeSet.Blue {
			minCubeSet.Blue = otherCubeSet.Blue
		}
		if otherCubeSet.Green > minCubeSet.Green {
			minCubeSet.Green = otherCubeSet.Green
		}
	}
	return minCubeSet
}

func isGamePossible(game Game) bool {
	for _, cubeSet := range game.CubeSets {
		if cubeSet.Red > 12 ||
			cubeSet.Green > 13 ||
			cubeSet.Blue > 14 {
			return false
		}
	}
	return true
}

func parseInputPartOne(inputFilePath string) ([]Game, error) {
	lines, err := utils.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, err
	}
	var games []Game
	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func parseGame(line string) (Game, error) {
	re := regexp.MustCompile("^Game ([0-9]+): (.*)$")
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return Game{}, fmt.Errorf("could not parse game: %s", line)
	}
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return Game{}, fmt.Errorf("could not parse game ID: %s", matches[1])
	}

	var cubeSets []CubeSet
	cubeSetStrings := strings.Split(matches[2], "; ")
	for _, cubeSetString := range cubeSetStrings {
		cubeSet, err := parseCubeSet(cubeSetString)
		if err != nil {
			return Game{}, err
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	return Game{
		ID:       id,
		CubeSets: cubeSets,
	}, nil
}

func parseCubeSet(s string) (CubeSet, error) {
	numAndColorStrings := strings.Split(s, ", ") // e.g. "5 red"
	cubeSet := CubeSet{}
	re := regexp.MustCompile("^([0-9]+) (.*)$")
	for _, numAndColorString := range numAndColorStrings {
		matches := re.FindStringSubmatch(numAndColorString)
		if matches == nil {
			return CubeSet{}, fmt.Errorf("could not parse number and colo '%s'", numAndColorString)
		}
		numCubes, err := strconv.Atoi(matches[1])
		if err != nil {
			return CubeSet{}, err
		}
		color := matches[2]
		switch color {
		case "red":
			cubeSet.Red = numCubes
		case "blue":
			cubeSet.Blue = numCubes
		case "green":
			cubeSet.Green = numCubes
		default:
			return CubeSet{}, fmt.Errorf("could not parse color: %s", color)
		}
	}
	return cubeSet, nil
}

type CubeSet struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	ID       int
	CubeSets []CubeSet
}
