package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Pairing struct {
	P1   int
	P2   int
	Area int
}

type ByArea []Pairing

func (a ByArea) Len() int           { return len(a) }
func (a ByArea) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByArea) Less(i, j int) bool { return a[i].Area < a[j].Area }

func area(p1 Point, p2 Point) int {
	var biggestX, smallestX = p1.X, p2.X
	var biggestY, smallestY = p1.Y, p2.Y

	if p2.X > p1.X {
		biggestX, smallestX = smallestX, biggestX
	}
	if p2.Y > p1.Y {
		biggestY, smallestY = smallestY, biggestY
	}

	return (biggestX - smallestX + 1) * (biggestY - smallestY + 1)
} 

func main() {
	fmt.Println("day 9 let's go!")

	filename := "day9.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day9Test.txt"

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

	var data []Point

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		data = append(data, Point{x, y})
	}

	var pairings []Pairing
	for idx, point := range data {
		for i := idx + 1; i < len(data); i++ {
			pairings = append(pairings, Pairing{idx, i, area(point, data[i])})
		}
	}
	sort.Sort(ByArea(pairings))

	fmt.Println(pairings)
}
