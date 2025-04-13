package finder

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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
	lenNeed := len(needText)

	selected := "\033[31m" + string(needText) + "\033[0m"
	for fileScanner.Scan() {
		line++
		res = fileScanner.Bytes()
		resOrig := res

		//if i := bytes.Index(res, needText); i > -1 {
		//	fmt.Println(strconv.Itoa(line) + ":" + strconv.Itoa(i+1) + "|" + string(res))
		//}
		deltaPos := 0
		finded := 0
		prefix := ""
		for x, d := bytes.Index(res, needText), 0; x > -1; x, d = bytes.Index(res, needText), d+x+1 {

			if finded == 0 {

				prefix += strconv.Itoa(line) + ":"
			}
			finded++
			deltaPos += x + lenNeed
			res = res[x+lenNeed:]
		}
		if finded > 0 {
			resSelected := bytes.ReplaceAll(resOrig, needText, []byte(selected))
			out = append(out, prefix+strconv.Itoa(finded)+"|"+string(resSelected))
		}
	}

	return &out
}
