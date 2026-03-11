package main

import (
	"fmt"

	"github.com/mY9Yd2/sgf"
)

func main() {
	root := sgf.LoadArgOrQuit(1)
	fmt.Println(root.SGF())
}
