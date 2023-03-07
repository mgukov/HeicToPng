package main

import (
	"os"

	"github.com/mgukov/HeicToPng/convert"
)

func main() {
	convert.ProcessArgs(os.Args[1:])
}
