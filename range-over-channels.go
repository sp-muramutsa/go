package main

import "fmt"

func main() {

	queue := make(chan string, 2)

	queue <- "value 1"
	queue <- "value 2"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}