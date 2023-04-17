package ex_1_saying_hello

import "fmt"
import "math/rand"

func Example1() {
	var firstName string
	greetings := []string{
		"Hello",
		"Howdy",
		"How's it going",
		"Hi",
	}

	fmt.Println("What is you name?")
	_, err := fmt.Scanln(&firstName)
	if err != nil {
		return
	}

	fmt.Printf("%s, %s, nice to meet you! ", greetings[rand.Intn(len(greetings))], firstName)
}
