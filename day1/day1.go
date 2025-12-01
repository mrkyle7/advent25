package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 1 - Let's go!")

	filename := "day1.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day1Test.txt"
	}
	fmt.Println("Opening file:", filename)

	file, err := os.Open(filename)

	
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Process the file

	scanner := bufio.NewScanner(file)
	var numZeros int64 = 0
	var timesPassedZero int64 = 0
	var dialNumber int64 = 50

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		direction := line[0]
		distancestr := line[1:]

		distance, err := strconv.ParseInt(distancestr, 10, 64)
		if err != nil {
			fmt.Println("Error converting distance:", err)
			return
		}

		numToMove := distance % 100
		fullCircuits := distance / 100
		fmt.Println("distance:", distance)
		fmt.Println("num to move:", numToMove)
		fmt.Println("full circuits:", fullCircuits)

		originalDialNumber := dialNumber
		switch direction {
		case 'L':
			dialNumber = dialNumber - numToMove
			if dialNumber < 0 {
				dialNumber = 100 + dialNumber
				if originalDialNumber != 0 {
					timesPassedZero++
				}
			}
			if dialNumber == 0 {
				timesPassedZero++
			}
		case 'R':
			dialNumber = dialNumber + numToMove
			if dialNumber >= 100 {
				dialNumber = dialNumber - 100
				timesPassedZero++
			}
		}

		timesPassedZero += fullCircuits

		fmt.Println("new dialing number:", dialNumber)

		fmt.Println("numZeros:", numZeros, "timesPassedZero:", timesPassedZero)
		fmt.Println("-----")
	}

	fmt.Println("answer:", "numZeros:", numZeros, "timesPassedZero:", timesPassedZero)
}
