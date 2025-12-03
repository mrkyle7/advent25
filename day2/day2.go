package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitByN(s string, n int) []string {
	var parts []string
	if len(s)%n != 0 {
		return parts
	}

	i := 0
	for range len(s) / n {
		parts = append(parts, s[i:i+n])
		i += n
	}
	return parts
}

func main() {
	fmt.Println("Day 2 - Let's go!")
	filename := "day2.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day2Test.txt"
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
		ranges := strings.Split(line, ",")
		fmt.Printf("Ranges: %q\n", ranges)
		for _, r := range ranges {
			if r == "" {
				continue
			}
			bounds := strings.Split(r, "-")
			fmt.Printf("Bounds: %q\n", bounds)
			var lower, upper int64
			lower, _ = strconv.ParseInt(bounds[0], 10, 64)
			upper, _ = strconv.ParseInt(bounds[1], 10, 64)
			fmt.Println("Lower:", lower, "Upper:", upper)
			for i := lower; i <= upper; i++ {
				// fmt.Println("number to inspect:", i)
				strI := strconv.Itoa(int(i))
				// if strLength % 2 == 0 {
				// 	firstHalf := strI[0:strLength / 2]
				// 	secondHalf := strI[strLength / 2:]
				// 	if firstHalf == secondHalf {
				// 		fmt.Println("first:", firstHalf, "second:", secondHalf)
				// 		total += int(i)
				// 	}
				// }
				valid := false
				for i := 1; i <= len(strI)/2; i++ {
					splitted := splitByN(strI, i)
					fmt.Printf("%q\n", splitted)

					if len(splitted) == 0 {
						continue
					}
					test := splitted[0]
					allMatch := true
					for _, i := range splitted {
						if i != test {
							allMatch = false
							break
						}
					}
					if allMatch {
						valid = true
						fmt.Println("found a match")
						break
					}
				}
				if valid {
					total += int(i)
				}
			}
		}
	}

	fmt.Println("Answer: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
