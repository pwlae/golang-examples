// Good article about reader https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

/*
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		if r+13 > 'Z' {
			return (r + 12) - 'Z' + 'A'
		}
		return r + 13

	} else if r >= 'a' && r <= 'z' {
		if r+13 > 'z' {
			return (r + 12) - 'z' + 'a'
		}
		return r + 13
	}
	return r
}

func (r *rot13Reader) Read(s []byte) (int, error) {
	size, err := r.r.Read(s)

	if err != nil {
		return size, err
	}

	var buffer []byte
	for i := 0; i < size; i++ {
		buffer = append(buffer, rot13(s[i]))
	}
	copy(s, buffer)
	return size, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
