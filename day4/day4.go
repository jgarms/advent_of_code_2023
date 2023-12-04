package main

import "fmt"
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
	return 0, nil
}

type Card struct {
	WinningNumbers map[int]bool
	PlayingNumbers map[int]bool
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

}
