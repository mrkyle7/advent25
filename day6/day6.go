package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var mathsWork [][]string

	// Your code logic here

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\s+`)
		cols := re.Split(line, -1)
		var maths []string
		for _, c := range cols {
			if c != "" {
				maths = append(maths, c)
			}
		}
		mathsWork = append(mathsWork, maths)

	}

	numWork := len(mathsWork[0])
	numNumbers := len(mathsWork) - 1

	for i:=0 ; i < numWork ; i++ {
		workResult := 0
		fmt.Println(i)
		switch mathsWork[numNumbers][i] {
		case "*":
			fmt.Println("Calculating * for", mathsWork[0][i])
			workResult = 1
			for num := 0; num < numNumbers; num++{
				calcNum, _ := strconv.Atoi(mathsWork[num][i])
				workResult *= calcNum
			}
		case "+":
			fmt.Println("Calculating + for", mathsWork[0][i])
			for num := 0; num < numNumbers; num++{
				calcNum, _ := strconv.Atoi(mathsWork[num][i])
				workResult += calcNum
			}
		}
		fmt.Println(workResult)
		total += workResult
	}
	fmt.Printf("%q\n", mathsWork)
	fmt.Println("Total", total)
}
