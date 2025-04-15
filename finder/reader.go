package finder

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Colour string

var (
	Reset   = Colour("\033[0m")
	Red     = Colour("\033[31m")
	Green   = Colour("\033[32m")
	Yellow  = Colour("\033[33m")
	Blue    = Colour("\033[34m")
	Magenta = Colour("\033[35m")
	Cyan    = Colour("\033[36m")
	Gray    = Colour("\033[37m")
	White   = Colour("\033[97m")
)

func ReadFromStdIn(needText []byte, c Colour) *[]string {
	reader := bufio.NewReader(os.Stdin)
	return ReadFromReaderLine(reader, needText, c)
}

func ReadFromFileLine(name string, needText []byte, c Colour) *[]string {
	reader, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer reader.Close()

	return ReadFromReaderLine(reader, needText, c)
}

func ReadFromReaderLine(reader io.Reader, needText []byte, c Colour) *[]string {

	out := []string{}

	fileScanner := bufio.NewScanner(reader)

	fileScanner.Split(bufio.ScanLines)

	line := 0
	selected := string(c) + string(needText) + string(Reset)
	for fileScanner.Scan() {
		line++
		res := fileScanner.Bytes()

		prefix := ""
		if x := bytes.Count(res, needText); x > 0 {
			prefix += strconv.Itoa(line) + ":" + strconv.Itoa(x)
			resSelected := bytes.Replace(res, needText, []byte(selected), x)
			out = append(out, prefix+"|"+string(resSelected))
		}
	}

	return &out
}
