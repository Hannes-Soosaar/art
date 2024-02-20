# CLI Art Decoder/Encoder

Art Decoder is a CLI tool to decode and encode ASCI art by compressing repeating string sections up to three characters long into the [numberOfReps repeatingSting] where numberOfReps and repeatingString is separated by as space.

## Starting The App

The main file is located in the cmd/art directory where you can use the "go run ." command to start the app

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

**$ go run . -e "myString to Decode"**

To enable multiline Encoding the first argument passed must be -m the second -e and the final argument must the file path.

**$ go run . -m  -e path/to/my/file/myFileToEncode.txt**

## Disable write to file

By default the program creates a output file for any multiline project that get encoded or decoded under "/assets/output/decoded" or "/assets/output/encoded respectively.

To disable creating encoded and decoded files. Change "WRITE_TO_FILE" value to false in "/constants/art_constants.go"

## Examples & Testing

The program includes a script that will run examples of each possible case and some common errors
to run the script:

Gives permission to launch the script file run in the terminal

**$ chmod +x run_cli.sh**

run the script while in the project directory with the following command.

**$ ./run_cli.sh**

this will run 30 instances of the app with various test cases. The run_cli.sh file arguments can be modified by editing the input arguments in run_cli.sh

**Running the script file clean up all of the text files from the "./assets/output/decoded" and *"./assets/output/decoded"* folders and cleans up the log file in *"./log/log.txt"*.**

# Web Art Decoder/Encoder 

Art Decoder/Encoder has a web interface to use the web graphical interface you must have Go installed your pc. The instruction are for running the program from the the linux terminal. If you are running a WSL on windows and port forwarding is enabled you will be able to access the interface from your windows browser.

## starting the Web app

To start the web app

1) permission needs to be granted to the launcher by running.

      **$ chmod +x run_web.sh**

2) starting the server  
   **$ ./run_web.sh**

  *the server can also be launched directly from the go file. the "main.go" file is located at ~/art/cmd/server/ to start the server run ~/art/cmd/server/$go run .*

## accessing the web App

 to access the web app open your standard a browser and navigate to  **localhost:8081** on the machine that you started the sever on.

### change listening port

to modify the port that the server is listening the file in *~/art/conf/config.go* can be modified if needed.

## Errors

In case of a error in returning a value for a decode or encode request a HTTP 400 Error will be sent. All other errors are standard.

## Contact

Have questions, concerns or issues related to the itinerary program pleas contact Hannes at  hannes.study@gmail.com

bugs should be report to hannes.study@gmail.com