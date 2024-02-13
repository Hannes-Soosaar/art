# art Decoder

Art Decoder is a CLI tool to decode and encode ASCI art by compressing repeating string sections up to three characters long into the [numberOfReps repeatingSting] where numberOfReps and repeatingString is separated by as space.

## Starting The App
 
The main file is located in the cmd/art directory where you can use the "go run ." command to start the app. 

 $ projectDirectory/cmd/art/ go run . -arg1 -arg2 -arg3

## Decoding

By default the program is in decode mode.

A single line passed in with quotation marks should be passed in as an argument.

**$ go run . "string to Decode"**

To Decode multiple lines the multiline decode mode must be toggled by passing in -m as the first arg and the path to the file containing multiple lines to decode after.

**$ go run . -m path/to/my/file/myFileToDecode.txt**

## Encoding

*The text to be encoded can not contain "[" or "]"*

To enable single line Encoding the first argument passed must be  -e  followed by the "text to Encode" in double quotes

** **$ go run . -e "myString to Decode"**

To enable multiline Encoding the first argument passed must be -m the second -e and the final argument must the file path.

**$ go run . -m  -e path/to/my/file/myFileToEncode.txt**

## Disable write to file

By default the program creates a output file for any multiline project that get encoded or decoded under "/assets/output/decoded" or "/assets/output/encoded respectively.

To disable creating encoded and decoded files. Change "WRITE_TO_FILE" value to false in "/constants/art_constants.go" 

## Examples & Testing

The program includes a script that will run examples of each possible case and some common errors
to run the script:

Gives permission to launch the script file run in the terminal

  **$ chmod +x run_art.sh**

run the script while in the project directory with the following command.

**$ ./run_art.sh**

this will run 30 instances of the app with various test cases. The run_art.sh file arguments can be modified by editing the input arguments in run_art.sh

*Running the sript file clean up all of the text files from the "./assets/output/decoded" and "./assets/output/decoded" folders and cleans up the log file in "./log/log.txt"*

## Contact 

Have questions, concerns or issues related to the itinerary program pleas contact Hannes at  hannes.study@gmail.com

bugs should be report to hannes.study@gmail.com