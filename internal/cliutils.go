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
