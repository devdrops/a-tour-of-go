package main

import (
	// Don't forget to `go get golang.org/x/tour/tree`
	"golang.org/x/tour/tree"
)

func doTheWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	doTheWalk(t.Left, ch)
	ch <- t.Value
	doTheWalk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	doTheWalk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
}

func main() {
}
