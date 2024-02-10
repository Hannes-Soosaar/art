package internal

import (
	"bufio"
	"os"

	"gitea.kood.tech/hannessoosaar/art/errors"
)

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		errors.ErrOpeningFile()
		panic(err)
	}
	return file
}

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
	return fileContent
}

func CreateFile(fileContent []string) {
	//TODO: add if time permits
}
