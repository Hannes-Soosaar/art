package internal

import (
	"fmt"
)

func DisplayDecodedSingle(decodedLine string) {
	fmt.Println(decodedLine)
}
func DisplayDecodedFile(fileContent []string) {
	for _, content := range fileContent {
		fmt.Printf("%s", content)
	}
}

func DisplayEncodedSingle(encodedLine string) {
	fmt.Println(encodedLine)
}

func DisplayEncodedFile(fileContent []string) {
	for _, content := range fileContent {
		fmt.Printf("%s", content)
	}
}
