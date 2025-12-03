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

		var bits []int64
		var currentIdx = 0
		for x:=1; x<=12; x++ {
			var currentMax int64 = 0
			for i := currentIdx; i < len(line)-12+x; i++ {
				var num, _ = strconv.ParseInt(string(line[i]), 10, 64)
				if num > currentMax {
					currentMax = num
					currentIdx = i + 1
				}
				if num == 9 {
					break
				}
			}
			fmt.Println("Current Max:", currentMax, "at index", currentIdx-1)
			bits = append(bits, currentMax)
		}	

		fmt.Println("Bits:", bits)

		var resultStr string
		for _, b := range bits {
			resultStr += strconv.FormatInt(b, 10)
		}
		fmt.Println("Result Str:", resultStr)
		var intResult, _ = strconv.ParseInt(resultStr, 10, 64)
		total += int(intResult)
	}
	fmt.Println("Total:", total)
}
