package main

import (
	"fmt"
	// "slices"
)

func main()	{
	fmt.Println("Hello, World!")
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


	whatAmiI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I am something else")
		}
	}

	whatAmiI(i)
	whatAmiI("hello")
	whatAmiI(3.14)

	var a [5]int
	fmt.Println("Array empty:", a)
	a[4] = 100
	fmt.Println("Array set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array b:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("Array b:", b)	

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d array:", twoD)

	var s []string
	fmt.Println("Slice empty:", s, s==nil, len(s)==0)

	s = make([]string, 3)
	fmt.Println("Slice initialized:", s, "cap:", cap(s), "len:", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Slice set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Slice appended:", s, "cap:", cap(s), "len:", len(s))
}