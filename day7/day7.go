package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 7 - Let's go!")

	filename := "day7.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day7Test.txt"

	}

	fmt.Println("Opening file:", filename)

	// Process the file

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, r := range line {
			switch string(r){
			case ".":
				row = append(row, 0)
			case "S":
				row = append(row, 1)
			case "^":
				row = append(row, -1)
			}
			
		}
		data = append(data, row)
	}

	total := 0

	rowLength := len(data[0])

	for rowIdx, row := range data {
		for colIdx, location := range row {

			// process first row differently
			if rowIdx == 0 {
				continue
			}
			// process . (0) with a beam above them
			if location == 0 && data[rowIdx-1][colIdx] > 0 {
				data[rowIdx][colIdx] = 1
				continue
			}
			// process splitters: ^ (-1)
			if location == -1 && data[rowIdx-1][colIdx] > 0 {
				//we split it!
				total++
				if colIdx != 0 {
					data[rowIdx][colIdx-1] = 1
				}
				if colIdx != rowLength {
					data[rowIdx][colIdx+1] = 1
				}
			}
		}
	}

	fmt.Println("updated data")

	for _, line := range data {
		fmt.Println(line)
	}
	fmt.Println("------§-------")
	fmt.Println("total", total)
}
