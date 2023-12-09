package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
import "advent_of_code_2023/utils"

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

	//result, err = partTwo("sample_input.txt")
	//if err != nil {
	//	fmt.Println("ERROR:", err)
	//	return
	//}
	//fmt.Println("Part two sample result:", result)
	//
	//result, err = partTwo("input.txt")
	//if err != nil {
	//	fmt.Println("ERROR:", err)
	//	return
	//}
	//fmt.Println("Part two result:", result)
}

func partOne(inputFilePath string) (int, error) {
	lines, err := utils.ReadFileLines(inputFilePath)
	if err != nil {
		return 0, err
	}
	cards, err := ParseCards(lines)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, card := range cards {
		values := card.GetMatchingValues()
		if len(values) == 0 {
			continue
		}
		valueToAdd := 1 << (len(values) - 1)
		sum += valueToAdd
	}
	return sum, nil
}

type Card struct {
	WinningNumbers []int
	PlayingNumbers []int
}

func (c Card) GetMatchingValues() []int {
	matches := make([]int, 0)
	for _, winningNumber := range c.WinningNumbers {
		for _, playingNumber := range c.PlayingNumbers {
			if winningNumber == playingNumber {
				matches = append(matches, winningNumber)
			}
		}
	}
	return matches
}

func ParseCards(lines []string) ([]Card, error) {
	cards := make([]Card, 0)
	for _, line := range lines {
		card, err := parseCard(line)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func parseCard(line string) (Card, error) {
	re := regexp.MustCompile("^Card\\s+\\d+:\\s+(.*)$")
	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		return Card{}, fmt.Errorf("could not parse card from '%s'", line)
	}
	allNumbers := matches[1]
	winnersAndPlaying := strings.Split(allNumbers, " | ")
	if len(winnersAndPlaying) != 2 {
		return Card{}, fmt.Errorf("could not find two groups of numbers for '%s'", allNumbers)
	}
	winningNumbers, err := parseNumbers(winnersAndPlaying[0])
	if err != nil {
		return Card{}, fmt.Errorf("could not parse numbers '%s'", winnersAndPlaying[0])
	}
	playingNumbers, err := parseNumbers(winnersAndPlaying[1])
	if err != nil {
		return Card{}, fmt.Errorf("could not parse numbers '%s'", winnersAndPlaying[1])
	}
	return Card{
		WinningNumbers: winningNumbers,
		PlayingNumbers: playingNumbers,
	}, nil
}

func parseNumbers(numString string) ([]int, error) {
	numStrings := strings.Fields(numString)
	nums := make([]int, 0)
	for _, s := range numStrings {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, val)
	}
	return nums, nil
}
