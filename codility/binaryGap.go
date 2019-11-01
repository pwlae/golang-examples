// Soulution doesn't work with large nums
package main

import "fmt"

func Solution(N int) int {
	result := 0
	temp_result := 0
	Nbinary := []rune(fmt.Sprintf("%b", N))

	for i := 0; i < len(Nbinary); i++ {
		if Nbinary[i] == rune('0') {
			temp_result += 1
		} else if temp_result > result {
			result = temp_result
			temp_result = 0
		}

	}
	return result
}

func main() {
	print(Solution(15))
}
