
#checks if go is installed
if command -v go &> /dev/null

then 
rm ./log/log.txt
# Array of arguments the number of arguments is the number of times the programm will be launced
arguments=(
    # "-h"
    # "-encode"
    # "-multi filetodecode.txt "
    # "-multi -encode filetoencode.txt "
    # " random text to see if it will triger help"
    "a[53][343][a32]" # true // does not work becuase of spaces 
    "asd][5s][gf]" #false
    "[adf]fd23]545[afd3253]"   #false 
    "[23][343]342]3423"   #false 
    "[[[[]]]]"  #true
    '"[2 3][3 4 3]3     42]3   423"'
)

# if the script is in the same folder, as the main go programm input "." 
program_path="./cmd/art/"

# Output file path and file name
output_file="./log/log.txt"


# ANSI escape code for red text
RED='\033[91m'
BLUE='\033[94m'
GREEN='\033[92m'
YELLOW='\033[93m'
# ANSI escape code to reset text color
RESET='\033[0m'


# Run the Go program for each set of arguments and save output to a file
for args in "${arguments[@]}"; do
    
    echo -e "${YELLOW}"running: $args
    echo ""
    echo -e "${GREEN}--------START--------${BLUE}"
    go run $program_path $args | tee -a $output_file
    # Add a sleep between runs if needed
    # sleep 1
    echo -e "${RED}-------END----------${REST}"
    echo ""
done

else 
echo "you need to install go, before running the app"
fi