package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
				freshIDs[low] = high
			}
			continue
		}

		if line == "" {
			continue
		}

		id, _ := strconv.Atoi(line)

		fresh := false

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

	totalFreshIds := 0
	sortedLows := []int{}

	for start := range freshIDs {
		sortedLows = append(sortedLows, start)
	}
	sort.Ints(sortedLows)

	currentMax := 0
	for i := 0; i < len(sortedLows); i++ {
		//3-5
		//10-14
		//11-14
		//12-13
		//13-18
		//14-17
		//15-21
		start := sortedLows[i]
		end := freshIDs[start]
		fmt.Println("start", start, "end", end)
		if i == 0 {
			totalFreshIds = end - start + 1
			currentMax = end
			fmt.Println(totalFreshIds)
			continue
		}
		if end <= currentMax {
			continue
		}
		if start <= currentMax {
			totalFreshIds += end - currentMax
		} else {
			totalFreshIds += end - start + 1
		}
		currentMax = end
		fmt.Println(totalFreshIds)
	}

	fmt.Println("Final total:", total)
	fmt.Println("Final total fresh IDs:", totalFreshIds)
}
