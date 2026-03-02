package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 2
	m["k2"] = 2

	fmt.Println(m)
	fmt.Println(m["k2"])
	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	_, present := m["k3"]
	fmt.Println(present)

	n := map[string]int{"foo": 1, "bar": 2}

	n2 := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(maps.Equal(n, n2))

}