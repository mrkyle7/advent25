package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5 - Let's go!")
	filename := "day5.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day5Test.txt"
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

	freshIDs := make(map[int]int)

	// Your code logic here

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Process each line

		if strings.Contains(line, "-") {
			
			parts := strings.Split(line, "-")
			low, _ := strconv.Atoi(parts[0])
			high, _ := strconv.Atoi(parts[1])
			fmt.Println("Adding range:", low, "to", high)
			if freshIDs[low] < high {
				fmt.Println("motherfucker", low)
				freshIDs[low] = high
			}
			continue
		}

		if line == "" {
			continue
		}

		id, _ := strconv.Atoi(line)

		fresh:= false

		for start, end := range freshIDs {
			if id >= start && id <= end {
				fmt.Println("ID", id, "is fresh (in range", start, "to", end, ")")
				total++
				fresh = true
				break
			}
		}

		if !fresh {
			fmt.Println("ID", id, "is not fresh")
		}
		
	}

	fmt.Println("Final total:", total)
}
