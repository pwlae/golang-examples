package main

import "fmt"

func reverseString(s []byte) {
    if len(s) <= 1 {
		return
	}

	tmp := s[0]
	s[0] = s[len(s)-1]
	s[len(s)-1] = tmp
	reverseString(s[1:len(s)-1])
}

func main() {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(str)
}