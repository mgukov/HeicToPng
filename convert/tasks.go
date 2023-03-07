package convert

import "fmt"

type ConvertTask struct {
	inputFile  string
	outputFile string
}

func (task ConvertTask) process() error {
	fmt.Printf("Converting %s to %s\n", task.inputFile, task.outputFile)
	return ConvertHeicToJpg(task.inputFile, task.outputFile)
}
