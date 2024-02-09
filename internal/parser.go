package internal

// Extracts sections correctly start with "[" and end with "]"
func ExtractEncodedSection(inputFromBracket string) (encodedSection string) {
	for i, char := range inputFromBracket {
		if char == ']' {
			return inputFromBracket[:i+1]
		}
	}
	return "BAD INPUT!"
}

// Removes the encoding markers
func RemoveFirstAndLastChar(encodedSection string) string {
	return encodedSection[1 : len(encodedSection)-1]
}
