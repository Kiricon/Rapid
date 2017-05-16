package templating

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
)

func AddPartial(file string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return insertPartial(FileString(cwd + "/" + file))
}

func FileString(file string) (string, string) {

	i := strings.LastIndex(file, "/")
	dir := file[0:i]
	filerc, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	return buf.String(), dir
}

func insertPartial(fileString string, dir string) string {
	r := regexp.MustCompile("{{\\s*[pP]artial\\s+(.*)\\s*}}")
	regArr := r.FindAllStringSubmatch(fileString, -1)

	for i := 0; i < len(regArr); i++ {
		path := strings.TrimSpace(dir + "/" + regArr[i][1])
		partial, _ := FileString(path)
		fileString = strings.Replace(fileString, regArr[i][0], partial, 1)
	}

	return fileString
}
