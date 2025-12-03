package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)


func main() {
	fmt.Println("Day 3 - Let's go!")
	filename := "day3.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day3Test.txt"
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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		
		// Your processing logic here

		var firstMax int64 = 0
		var secondMax int64 = 0
		var idx = 0
		for i:= 0; i < len(line) - 1; i++ {
			var num, _ = strconv.ParseInt(string(line[i]), 10, 64)
			if num > firstMax {
				firstMax = num
				idx = i
			}
			if (num == 9) {
				break
			}
		}
		for i:= idx + 1; i < len(line); i++ {
			var num, _ = strconv.ParseInt(string(line[i]), 10, 64)
			if num > secondMax {
				secondMax = num
			} 
			if num == 9 {
				break
			}
		}
		fmt.Println("Max:", firstMax, "at index", idx)
		fmt.Println("Second Max:", secondMax)
		var result = strconv.Itoa(int(firstMax)) + strconv.Itoa(int(secondMax))
		var intResult, _ = strconv.ParseInt(result, 10, 64)
		fmt.Println("Result for line:", result)
		total += int(intResult)

	}
	fmt.Println("Part 1 Total:", total)
}