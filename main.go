package main

import (
	"errors"
	"fmt"
	"grep/finder"
	"io/fs"
	"os"
	time "time"
)

var (
	needText = []byte{}
	Reset    = "\033[0m"
	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Blue     = "\033[34m"
	Magenta  = "\033[35m"
	Cyan     = "\033[36m"
	Gray     = "\033[37m"
	White    = "\033[97m"
)

type (
	color struct {
		Start int
		End   int
	}
	colorMap map[int][]color

	lineMap map[int]string
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	args := os.Args
	fmt.Println(args)
	file := "."
	if len(args) > 2 {
		file = args[len(args)-1]
	}
	if len(args) > 1 {
		needText = []byte(args[1])
	} else {
		return
	}

	timeStart := time.Now()
	if ok, err := IsDir(file); err == nil && ok {
		ReadFromDir(file, needText)
	} else if err == nil && !ok {
		ReadFromFile(file, needText)
	} else if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Sub(timeStart).String())

}

func ReadFromFile(filename string, needText []byte) {
	if filename == "" {
		return
	}
	for _, line := range *finder.ReadFromFileLine(filename, needText) {
		fmt.Println(line)
	}
	return
}

func ReadFromDir(dirname string, needText []byte) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, e := range entries {
		nameFile := e.Name()
		if ok, _ := IsDir(nameFile); !ok {
			fmt.Println(Green + nameFile + Reset)
			for _, line := range *finder.ReadFromFileLine(e.Name(), needText) {
				fmt.Println(line)
			}
		}

	}

	return nil
}

func IsDir(filename string) (bool, error) {

	s, err := os.Stat(filename)

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	if s.IsDir() {
		return true, nil
	}
	return false, err

}
