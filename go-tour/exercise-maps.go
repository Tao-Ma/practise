/*
 * exercise-maps.go
 */

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)

	for _, w := range strings.Split(s, " ") {
		if _, ok := m[w]; ok {
			m[w] = m[w] + 1
		} else {
			m[w] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}
