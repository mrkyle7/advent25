package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 6 - Let's go!")
	filename := "day6.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day6Test.txt"
	}
	fmt.Println("Opening file:", filename)

	// Process the file

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	// Your code logic here

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var colLengths []int
	currentCol := 0

	for idx, char := range lines[len(lines)-1] {
		s := string(char)
		if idx == 0 {
			continue
		}
		if s == " " {
			currentCol++
		} else {
			colLengths = append(colLengths, currentCol)
			currentCol = 0
		}
	}

	colLengths = append(colLengths, currentCol+1)
	fmt.Println("Col lengths", colLengths)

	var mathsWork [][]string

	for _, line := range lines {
		fmt.Printf("%q\n", line)
		currentIdx := 0
		var row []string
		for _, length := range colLengths {
			row = append(row, line[currentIdx:currentIdx+length])
			currentIdx += length + 1
		}
		mathsWork = append(mathsWork, row)
	}

	numWork := len(mathsWork[0])
	numNumbers := len(mathsWork) - 1

	var actualMaths [][]string

	for worki := range numWork {
		var realNumbers []string
		for bit := colLengths[worki]; bit > 0; bit-- {
			var newNumString string
			for num := range numNumbers {
				fmt.Printf("dealing with %q\n", mathsWork[num][worki])
				strBit := string(mathsWork[num][worki][bit-1])
				if strBit != " " {
					newNumString += strBit
				}
			}
			fmt.Println(newNumString)
			realNumbers = append(realNumbers, newNumString)
		}
		realNumbers = append(realNumbers, string(mathsWork[numNumbers][worki][0]))
		fmt.Println(realNumbers)
		actualMaths = append(actualMaths, realNumbers)

	}

	for _, maths := range actualMaths {
		fmt.Println("doing maths for:", maths)
		switch maths[len(maths)-1] {
		case "*":
			result := 1
			for x := 0; x < len(maths)-1; x++ {
				num, _ := strconv.Atoi(maths[x])
				result *= num
			}
			total += result
		case "+":
			result := 0
			for x := 0; x < len(maths)-1; x++ {
				num, _ := strconv.Atoi(maths[x])
				result += num
			}
			total += result
		}
	}

	fmt.Printf("%q\n", mathsWork)
	fmt.Println(actualMaths)
	fmt.Println("Total", total)
}
