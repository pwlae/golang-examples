package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		f := 0

		if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		} else {
			a := 0
			b := 1

			for i:=1; i < x; i++ {
				f = a + b
				a = b
				b = f
			}
		}
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
