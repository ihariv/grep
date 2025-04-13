package finder

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type color string

var (
	needText = []byte{}
	Reset    = color("\033[0m")
	Red      = color("\033[31m")
	Green    = color("\033[32m")
	Yellow   = color("\033[33m")
	Blue     = color("\033[34m")
	Magenta  = color("\033[35m")
	Cyan     = color("\033[36m")
	Gray     = color("\033[37m")
	White    = color("\033[97m")
)

func ReadFromFileLine(name string, needText []byte) *[]string {

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

	selected := "\033[31m" + string(needText) + "\033[0m"
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
