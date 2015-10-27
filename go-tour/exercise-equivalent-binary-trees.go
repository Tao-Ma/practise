package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) (r bool) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(t *tree.Tree, ch chan int) { Walk(t, ch); close(ch) }(t1, ch1)
	go func(t *tree.Tree, ch chan int) { Walk(t, ch); close(ch) }(t2, ch2)

	r = true
	for i := range ch1 {
		select {
		case j := <-ch2:
			if i != j {
				r = false
			}
		}
	}

	return
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	go func() { Walk(t, ch); close(ch) }()
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
