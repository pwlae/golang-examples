// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(s []byte) (int, error) {
	for i := 0; i < len(s); i++ {
		s[i] = 'A'
	}
	return len(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
