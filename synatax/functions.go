package main

import (
	"fmt"
)

// Return function
func showTextFunction(text string) func() string {
	return func() string {
		return text + " function"
	}
  }

// Return string
func showTestString(text string) string {
	return func() string {
		return text + " string"
	}()
}

// Return function
func showTestPuzzled() func(string) string {
	postfix := func() string {
		return " puzzled"
	}()

	return func(text string) string {
		return "I'm " + text + postfix 
	}
}

func main() {
	text1 := showTextFunction("yo yo yo")
	fmt.Println(text1())

	text2 := showTestString("yo yo yo")
	fmt.Println(text2)

	text3 := showTestPuzzled()
	fmt.Println(text3("very"))
}