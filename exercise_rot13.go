package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i, val := range b[:n] {
		switch {
			case val >= 'A' && val <= 'Z':
				b[i] = byte(((val-'A'+13) % 26)+'A')
			case val >= 'a' && val <= 'z':
				b[i] = byte(((val-'a'+13) % 26)+'a')
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}