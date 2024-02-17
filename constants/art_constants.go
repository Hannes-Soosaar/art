package constants

var (
	HELP                  = "To use the decoder: \n\n for help\n $go run . -h \n\n Decode a single line run: \n $go run . \"string to decode\" \n\n Decode multiple lines run:\n $go run . -m path/to/my/file/myFileToDecode.txt \n\n Encode a single line run: \n $go run . -e \"string to encode\" \n\n Encode multiple line run: \n $go run . -m  -e path/to/my/file/myFileToEncode.txt\n"
	ERR_READING_FILE      = "Error: \n reading file"
	ERR_WRITING_FILE      = "Error: \n writing file"
	ERR_ARG1_NOT_NUM      = "Error: \n The first argument is not a number"
	ERR_ARG_NO_SPACE      = "Error: \n The arguments are not separated by a space"
	ERR_ARG_NO_SECOND_ARG = "Error: \n There is no second argument"
	ERR_ARG_UNBALANCED    = "Error: \n Square brackets are unbalanced"
	ERR_OPENING_FILE      = "Error: \n file does not exist or the path is wrong"
	ERR_ENCODING_CHAR     = "Error: \n  can not encode input with '[' or ']' "

	WRITE_TO_FILE = true
)
