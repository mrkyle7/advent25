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
	Z int
}

type Pairing struct {
	P1       int
	P2       int
	Distance float64
}

type ByDistance []Pairing

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt((math.Pow(float64(p1.X)-float64(p2.X), 2)) + (math.Pow(float64(p1.Y)-float64(p2.Y), 2) + (math.Pow(float64(p1.Z)-float64(p2.Z), 2))))
}

func main() {
	fmt.Println("day 8 let's go!")

	filename := "day8.txt"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		filename = "day8Test.txt"

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
		z, _ := strconv.Atoi(split[2])
		data = append(data, Point{x, y, z})
	}

	var pairings []Pairing
	for idx, point := range data {
		for i := idx + 1; i < len(data); i++ {
			pairings = append(pairings, Pairing{idx, i, distance(point, data[i])})
		}
	}
	sort.Sort(ByDistance(pairings))

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
