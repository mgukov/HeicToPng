package main

import (
	"log"
	"os"
)

func main() {

	heic := os.Args[1]
	out := os.Args[2]

	if len(heic) == 0 {
		panic("Empty input file")
	}

	if len(out) == 0 {
		panic("Empty output file")
	}

	log.Println("Convert " + heic + " to " + out)

	err := ConvertHeicToJpg(heic, out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conversion Passed")
}
