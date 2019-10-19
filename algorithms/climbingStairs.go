package main

// Solution using factorial doesn't work when result is out of int range
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	
	return x * factorial(x-1)
}

func climbStairs(n int) int {
	c := 0
	for i := 0; i <= n; i++ {
		c += factorial(n)/(factorial(n-i)*factorial(i))
		n -= 1
	}
	return c
}

func main() {
	print(climbStairs(15))
}
