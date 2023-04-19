package ex_2_counting_num_chars

import (
	"fmt"
)

func Example() {
	var input string
	for true {
		fmt.Println("What is the input string?")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Input is required, try again.")
			continue
		} else {
			break
		}
	}

	fmt.Printf("%s has %d characters.", input, len(input))
}
