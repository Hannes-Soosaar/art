package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitea.kood.tech/hannessoosaar/art/errors"
)
// Opens a file and returns the  file object
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		errors.ErrOpeningFile()
		panic(err)
	}
	return file
}
// Reads the content of a file and returns the file as a slice of string
func GetFileContent(file *os.File) (fileContent []string) {
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}
	if err := scanner.Err(); err != nil {
		errors.ErrReadingFile()
		return nil
	}
	file.Close()
	return fileContent
}

func CreateEncodedFile(fileContent []string, filePath string) {
	fileName := strings.Split(filePath, "/") // split the string at "/"
	writePath := fileName[len(fileName)-1]   // gets the file name
	writePath = "./assets/output/encoded/encoded_" + writePath
	file, err := os.Create(writePath)
	if err != nil {
		errors.ErrWritingFile()
	}
	defer file.Close()
	for _, content := range fileContent {
		_, err = fmt.Fprint(file, content)
		if err != nil {
			errors.ErrWritingFile()
		}
	}
}
func CreateDecodedFile(fileContent []string, filePath string) {
	fileName := strings.Split(filePath, "/")
	writePath := fileName[len(fileName)-1]
	writePath = "./assets/output/decoded/decoded_" + writePath
	file, err := os.Create(writePath)
	if err != nil {
		errors.ErrWritingFile()
	}
	defer file.Close()
	for _, content := range fileContent {
		_, err = fmt.Fprint(file, content)
		if err != nil {
			errors.ErrWritingFile()
		}
	}
}
