package main

import (
	"bufio"
	"fmt"
	"math"
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

	return (biggestX - smallestX) * (biggestY - smallestY)
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

	var clusters []map[int]struct{}

	var i int

	var lastPair Pairing
	for len(clusters) == 0 || len(clusters[0]) != len(data) {
		fmt.Println(pairings[i])
		lastPair = pairings[i]
		addedToExisting := false
		if len(clusters) == 0 {
			nodes := make(map[int]struct{})
			nodes[pairings[i].P1] = struct{}{}
			nodes[pairings[i].P2] = struct{}{}
			clusters = append(clusters, nodes)
		} else {
			for c, cluster := range clusters {
				var toAdd []int
				if _, ok := cluster[pairings[i].P1]; ok {
					toAdd = append(toAdd, pairings[i].P2)
				}
				if _, ok := cluster[pairings[i].P2]; ok {
					toAdd = append(toAdd, pairings[i].P1)
				}

				for _, node := range toAdd {
					cluster[node] = struct{}{}
				}

				if len(toAdd) == 1 {
					addedToExisting = true
					for c1 := c + 1; c1 < len(clusters); c1++ {
						if _, ok := clusters[c1][toAdd[0]]; ok {
							for node := range clusters[c1] {
								cluster[node] = struct{}{}
							}
							clusters[c1] = make(map[int]struct{})
							break
						}
					}
				}
				if len(toAdd) == 2 {
					addedToExisting = true
					break
				}
			}
			if !addedToExisting {
				nodes := make(map[int]struct{})
				nodes[pairings[i].P1] = struct{}{}
				nodes[pairings[i].P2] = struct{}{}
				clusters = append(clusters, nodes)
			}
		}
		i++
	}

	fmt.Println(clusters)

	var lengths []int

	for _, c := range clusters {
		lengths = append(lengths, len(c))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(lengths)))

	fmt.Println(lengths)

	var total int = 1

	for i := range 3 {
		total *= lengths[i]
	}

	fmt.Println(total)

	fmt.Println(lastPair)
	fmt.Println(data[lastPair.P1].X * data[lastPair.P2].X)
}
