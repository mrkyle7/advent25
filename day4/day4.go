package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4 - Let's go!")
	filename := "day4.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day4Test.txt"
	}
	fmt.Println("Opening file:", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Process the file

	total := 0
	scanner := bufio.NewScanner(file)

	var rolls [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		rolls = append(rolls, row)
	}

	colLen := len(rolls[0])

	var ansRolls [][]string
	totalChanged := true
	for totalChanged {
		ansRolls = [][]string{}
		startTotal := total
		fmt.Println("checking, current total:", total)
		for row := 0; row < len(rolls); row++ {
			var ansRow []string
			for col := 0; col < len(rolls[row]); col++ {
				adjTotal := 0
				// fmt.Println(row, col, rolls[row][col])
				if rolls[row][col] != "@" {
					ansRow = append(ansRow, ".")
					// fmt.Println(".")
					continue
				}
				for adjRow := row - 1; adjRow <= row+1; adjRow++ {
					for adjCol := col - 1; adjCol <= col+1; adjCol++ {
						if adjCol < 0 || adjCol == colLen || adjRow < 0 || adjRow == len(rolls) || (adjCol == col && adjRow == row) {
							continue
						}
						if rolls[adjRow][adjCol] == "@" {
							// fmt.Println(adjRow, adjCol)
							adjTotal++
						}
					}
				}
				if adjTotal < 4 {
					total++
					// fmt.Println("x", adjTotal)
					ansRow = append(ansRow, "x")
				} else {
					// fmt.Println("@", adjTotal)
					ansRow = append(ansRow, "@")
				}
			}
			// fmt.Println(ansRow)
			ansRolls = append(ansRolls, ansRow)
		}
		rolls = ansRolls
		if total == startTotal {
			totalChanged = false
		}
	}

	fmt.Println(rolls)
	fmt.Println("---------------")
	fmt.Println(ansRolls)

	fmt.Println("Answer:", total)
}
