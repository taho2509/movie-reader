package interfaces

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type fileReader struct {
	FileName string
}

func (fr *fileReader) Read() []string {
	csvFile, _ := os.Open(fr.FileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var titles []string = []string{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		titles = append(titles, line[0])
	}
	return titles
}

// NewFileReader return a concrete implementation of a DataReader
func NewFileReader(filename string) *fileReader {
	return &fileReader{
		FileName: filename,
	}
}
