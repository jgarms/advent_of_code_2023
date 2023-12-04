package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strconv"
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

func partTwo(inputFilePath string) (int, error) {
	lines, err := utils.ReadFileLines(inputFilePath)
	if err != nil {
		return 0, err
	}
	schematic := CreateSchematic(lines)
	return schematic.getGearSum(), nil
}

func partOne(inputFilePath string) (int, error) {
	lines, err := utils.ReadFileLines(inputFilePath)
	if err != nil {
		return 0, err
	}
	schematic := CreateSchematic(lines)
	result := schematic.getPartSum()

	return result, nil
}

type Schematic struct {
	Contents           []string
	MaxX               int
	MaxY               int
	PartNumbers        []PartNumber
	GearsToPartNumbers map[Coordinate][]PartNumber
}

func CreateSchematic(lines []string) *Schematic {
	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	schematic := &Schematic{
		Contents:           lines,
		MaxX:               maxX,
		MaxY:               maxY,
		PartNumbers:        make([]PartNumber, 0),
		GearsToPartNumbers: make(map[Coordinate][]PartNumber),
	}
	schematic.parse()
	return schematic
}

func (schematic *Schematic) parse() {
	var currentPartNumber *PartNumber
	for y := 0; y <= schematic.MaxY; y++ {
		for x := 0; x <= schematic.MaxX; x++ {
			c := schematic.Contents[y][x]
			if isNum(c) {
				if currentPartNumber == nil {
					currentPartNumber = &PartNumber{
						X:     x,
						Y:     y,
						Value: "",
					}
				}
				currentPartNumber.Value += string(c)
			} else {
				if currentPartNumber != nil {
					// We've reached the end of a part number
					if currentPartNumber.IsValid(schematic) {
						schematic.PartNumbers = append(schematic.PartNumbers, *currentPartNumber)
					}
					coordinate, gearErr := currentPartNumber.GetGearCoordinate(schematic)
					if gearErr == nil {
						schematic.GearsToPartNumbers[coordinate] = append(schematic.GearsToPartNumbers[coordinate], *currentPartNumber)
					}
					currentPartNumber = nil
				}
			}
		}
	}
}

func (schematic *Schematic) getPartSum() int {
	sum := 0
	for _, partNumber := range schematic.PartNumbers {
		sum += partNumber.GetValue()
	}
	return sum
}

func (schematic *Schematic) getGearSum() int {
	sum := 0
	for _, partNumbers := range schematic.GearsToPartNumbers {
		if len(partNumbers) == 2 {
			product := partNumbers[0].GetValue() * partNumbers[1].GetValue()
			sum += product
		}
	}
	return sum
}

func (schematic *Schematic) GetCharAt(x, y int) (uint8, error) {
	if x < 0 || x > schematic.MaxX {
		return 0, fmt.Errorf("x is out of bounds: %d", x)
	}
	if y < 0 || y > schematic.MaxY {
		return 0, fmt.Errorf("y is out of bounds: %d", y)
	}
	return schematic.Contents[y][x], nil
}

func isNum(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c uint8) bool {
	if c == '.' {
		return false
	}
	if isNum(c) {
		return false
	}
	return true
}

func isGearSymbol(c uint8) bool {
	if c == '*' {
		return true
	}
	return false
}

type PartNumber struct {
	X     int
	Y     int
	Value string
}

func (partNumber PartNumber) IsValid(schematic *Schematic) bool {
	// A part number is valid if it has a symbol anywhere on its border
	for y := partNumber.Y - 1; y <= partNumber.Y+1; y++ {
		for x := partNumber.X - 1; x <= partNumber.X+len(partNumber.Value); x++ {
			c, err := schematic.GetCharAt(x, y)
			if err != nil {
				continue
			}
			if isSymbol(c) {
				return true
			}
		}
	}
	return false
}

func (partNumber PartNumber) GetGearCoordinate(schematic *Schematic) (Coordinate, error) {
	// If a part number has a gear symbol on its border, return it
	for y := partNumber.Y - 1; y <= partNumber.Y+1; y++ {
		for x := partNumber.X - 1; x <= partNumber.X+len(partNumber.Value); x++ {
			c, err := schematic.GetCharAt(x, y)
			if err != nil {
				continue
			}
			if isGearSymbol(c) {
				return Coordinate{
					X: x,
					Y: y,
				}, nil
			}
		}
	}
	return Coordinate{}, fmt.Errorf("No gear symbol found")
}

func (partNumber PartNumber) GetValue() int {
	value, err := strconv.Atoi(partNumber.Value)
	if err != nil {
		return 0
	}
	return value
}

type Coordinate struct {
	X int
	Y int
}
