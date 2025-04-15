package main

import (
	"fmt"
	"grep/finder"
	"os"
	"time"
)

var (
	needText = []byte{}
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

	args := os.Args[1:]
	fmt.Println(args)
	file := "."
	if len(args) > 0 {
		needText = []byte(args[0])
	} else {
		return
	}

	if len(args) > 1 {
		file = args[len(args)-1]
	} else {
		file = ""
	}

	timeStart := time.Now()
	if file == "" {
		ReadFromStdIn(needText)
	} else {
		if ok, err := IsDir(file); err == nil && ok != nil && *ok {
			ReadFromDir(file, needText)
		} else if err == nil && ok != nil && !*ok {
			ReadFromFile(file, needText)
		} else if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(time.Now().Sub(timeStart).String())

}

func ReadFromFile(filename string, needText []byte) {
	if filename == "" {
		return
	}
	for _, line := range *finder.ReadFromFileLine(filename, needText, finder.Blue) {
		fmt.Println(line)
	}
	return
}

func ReadFromStdIn(needText []byte) {

	for _, line := range *finder.ReadFromStdIn(needText, finder.Blue) {
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
		if ok, err := IsDir(nameFile); err == nil && ok != nil && !*ok {
			fmt.Println(string(finder.Green) + nameFile + string(finder.Reset))
			for _, line := range *finder.ReadFromFileLine(e.Name(), needText, finder.Red) {
				fmt.Println(line)
			}
		}

	}

	return nil
}

func IsDir(filename string) (*bool, error) {

	res := false
	s, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		res = true
	}
	return &res, err

}
