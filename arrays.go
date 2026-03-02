package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp: ", a)
	fmt.Println(len(a))

	a[4] = 34
	fmt.Println("set: ", a)
	fmt.Println("get at index 4", a[4])
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	b= [...]int{10, 20, 30, 40, 50}
	fmt.Println("dcl: ", b)

	b= [...]int{10, 3: 20, 50}
	fmt.Println("dcl: ", b)

	var twoD [2][3] int 
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	for i := range 2 {
		for j := range 3 {
			fmt.Print(twoD[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

	fmt.Print(twoD)


}