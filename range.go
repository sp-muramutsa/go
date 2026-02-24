package main

import "fmt" 

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for k, v := range m {
		fmt.Println("key: ", k, ", value: ", v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
} 