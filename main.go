package main

import (
	"github.com/jealous/fanhao/lib"
	"os"
)

func main() {
	args := os.Args
	var folder string
	if len(args) > 1 {
		folder = args[1]
	} else {
		folder = fanhao.CurrentFolder()
	}
	fanhao.NormalizeAll(folder)
}
