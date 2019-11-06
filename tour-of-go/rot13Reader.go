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
	if (r >= 'A' && r <= 'Z') {
		if r + 13 > 'Z' {
			return (r + 12) - 'Z' + 'A'
		}
		return r + 13

	} else if (r >= 'a' && r <= 'z') {
		if r + 13 > 'z' {
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
	for i:=0; i<size; i++ {
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
