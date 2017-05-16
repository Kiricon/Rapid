package templating

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
)

func AddPartial(file string) string {
	return insertPartial(FileString(file))
}

func FileString(file string) string {
	filerc, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	return buf.String()
}

func insertPartial(fileString string) string {
	r := regexp.MustCompile("{{\\s*[pP]artial\\s+(.*)\\s*}}")
	regArr := r.FindAllStringSubmatch(fileString, -1)

	for i := 0; i < len(regArr); i++ {
		partial := FileString(strings.TrimSpace(regArr[i][1]))
		fileString = strings.Replace(fileString, regArr[i][0], partial, 1)
	}

	return fileString
}
