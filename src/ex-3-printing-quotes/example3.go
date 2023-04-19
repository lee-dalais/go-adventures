package ex_3_printing_quotes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Example() {
	var inputQuote string
	var inputAuthor string
	var err error
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("What is the quote? ")
		inputQuote, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("A quote is required, try again.")
			continue
		}

		inputQuote = strings.TrimRight(inputQuote, "\n")
		break
	}

	for true {
		fmt.Print("Who said it? ")
		inputAuthor, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("An author is required, try again.")
			continue
		}

		inputAuthor = strings.TrimRight(inputAuthor, "\n")
		break
	}

	fmt.Printf("%s says \"%s\"", inputAuthor, inputQuote)
}
