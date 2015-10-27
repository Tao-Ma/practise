package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// r MyReader/r *MyReader are OK!
func (r MyReader) Read(b []byte) (int, error) {
	for idx := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
