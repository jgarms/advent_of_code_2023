package main

import (
	"bufio"
	"fmt"
	"os"
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
}

func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func partOne(inputFilePath string) (int, error) {
	lines, err := readFileLines(inputFilePath)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, line := range lines {
		lineValue := parseLinePartOne(line)
		fmt.Println(lineValue)
		total += lineValue
	}
	return total, nil
}

func parseLinePartOne(line string) int {
	firstNum := -1
	secondNum := -1
	for _, c := range line {
		num := parseCharToInt(c)
		if num > -1 {
			if firstNum == -1 {
				firstNum = num
			} else {
				secondNum = num
			}
		}
	}
	if firstNum == -1 {
		// we didn't find any digits
		return 0
	} else if secondNum == -1 {
		// No second number found, so use the first number twice
		secondNum = firstNum
	}
	return firstNum*10 + secondNum
}

func parseCharToInt(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}
	return -1
}
