package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	if _, err := rr.r.Read(b); err != nil {
		return 0, err
	}

	for i := range b {
		switch {
		case b[i] >= 'a' && b[i] <= 'm':
			fallthrough
		case b[i] >= 'A' && b[i] <= 'M':
			b[i] += 13
		case b[i] >= 'n' && b[i] <= 'z':
			fallthrough
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] -= 13
		default:
			continue
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
