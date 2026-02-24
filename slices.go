package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {

	var s []string
	fmt.Println("uninitialized string: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "u"
	s[1] = "m"
	s[2] = "u"

	fmt.Println("set: ", s)
	fmt.Print("get: ", s[2])

	s = append(s, "n")
	s = append(s, "t")
	s = append(s, "u")

	fmt.Println("after appending: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of s: ", c)

	fmt.Println(s[2:5])
	fmt.Println(s[:3])
	fmt.Println(s[3:])

	t := []string {"g", "h", "i"}
	t2 := []string {"g", "h", "i"}

	fmt.Println(slices.Equal(t, t2))

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}