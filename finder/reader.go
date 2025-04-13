package finder

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Colour string

var (
	needText = []byte{}
	Reset    = Colour("\033[0m")
	Red      = Colour("\033[31m")
	Green    = Colour("\033[32m")
	Yellow   = Colour("\033[33m")
	Blue     = Colour("\033[34m")
	Magenta  = Colour("\033[35m")
	Cyan     = Colour("\033[36m")
	Gray     = Colour("\033[37m")
	White    = Colour("\033[97m")
)

func ReadFromFileLine(name string, needText []byte, c Colour) *[]string {

	out := []string{}

	readFile, err := os.Open(name)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	res := []byte{}

	line := 0

	selected := string(c) + string(needText) + string(Reset)
	for fileScanner.Scan() {
		line++
		res = fileScanner.Bytes()

		prefix := ""
		if x := bytes.Count(res, needText); x > 0 {
			prefix += strconv.Itoa(line) + ":" + strconv.Itoa(x)
			resSelected := bytes.ReplaceAll(res, needText, []byte(selected))
			out = append(out, prefix+"|"+string(resSelected))
		}
	}

	return &out
}
