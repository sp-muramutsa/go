package main

import "fmt"

func zeroval(value int) {
	value = 0
}

func zeroptr(ptrValue *int) {
	*ptrValue = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("after passing by value: ", i)

	zeroptr(&i)
	fmt.Println("after passing by pointer: ", i)
	fmt.Println("Pointer: ", &i)

}
