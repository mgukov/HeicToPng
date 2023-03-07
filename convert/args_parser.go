package convert

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func checkArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("Invalid arguments")
	}

	first := args[0]

	if first == "-f" {
		if len(args) < 2 {
			return errors.New("Invalid arguments")
		}
	}

	return nil
}

func ProcessArgs(args []string) {
	if checkArgs(args) != nil {
		log.Fatalln("Usages: \n - for single file: heic2png in.hect [out.png]\n-for directory: heic2png -f in [out]")
	}

	first := args[0]

	if first == "-f" {
		err := processForlderModeArgs(args[1:])
		if err != nil {
			log.Fatalf("Unagle to convert images: %s", err.Error())
		}
	} else {
		err := processSingleFileModeArgs(args)

		if err != nil {
			log.Fatalf("Unagle to convert image: %s", err.Error())
		}
	}
}

func getOutFileName(name string) string {
	return name[:len(name)-len(filepath.Ext(name))] + ".png"
}

func processSingleFileModeArgs(args []string) error {

	in := args[0]
	out := getOutFileName(in)

	if len(args) > 1 {
		out = args[1]
	}

	task := ConvertTask{
		inputFile:  in,
		outputFile: out,
	}

	return task.process()
}

func processForlderModeArgs(args []string) error {

	in := args[0]
	out := in

	if len(args) > 1 {
		out = args[1]
	}

	err := os.MkdirAll(out, os.ModePerm)
	if err != nil {
		return err
	}

	err, tasks := getTasks(in, out)

	if err != nil {
		return err
	}

	for _, t := range tasks {
		err = t.process()
		if err != nil {
			fmt.Printf("Unable to convert %s to %s: %s\n", t.inputFile, t.outputFile, err.Error())
		}
	}

	return nil
}

func getTasks(input, output string) (err error, tasks []ConvertTask) {

	files, err := ioutil.ReadDir(input)
	if err != nil {
		return
	}

	tasks = []ConvertTask{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())

		sortName := filepath.Base(file.Name())
		if strings.EqualFold(ext, ".heic") {
			tasks = append(tasks, ConvertTask{input + "/" + file.Name(), output + "/" + getOutFileName(sortName)})
		}
	}

	return
}
