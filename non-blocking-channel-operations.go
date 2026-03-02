package main 

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case result := <-messages:
		fmt.Println(result)
	default:
		fmt.Println("No messages")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("Message sent")
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	case sig := <-signals:
		fmt.Println("Signal received", sig)
	default:
		fmt.Println("No activity")
	}
}