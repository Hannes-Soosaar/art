
#checks if go is installed
if command -v go &> /dev/null

then 
rm ./log/log.txt
# Array of arguments the number of arguments is the number of times the programm will be launced
arguments=(
    # "-h"
    # "-encode"
    # "-multi filetodecode.txt"
    # "-multi -encode filetoencode.txt "
    # " random text to see if it will triger help"
    #  "-h"  # test help / PASS
    # ' -encode "[4 5][3 34]"' # test encode // PASS

)

sl_arguments=(
    "[5 3][3 3][3 2]" # true balanced brackets no errors // PASS
    "[5 3][3 3][3 2]" # true balanced brackets no errors // PASS
    # "asd][5 s][gf]" #false unbalance brackets // PASS
    # '"[a df]fd23545[afd3253]"'   #false first arg not a number // PASS
    "[15 #][5 -_]-[5 #]"   #false second arg a space // PASS
    "[19 D]"   #false second arg a space // PASS
    #  '"[2    d"'   #false unbalanced // Pass
    #false unbalanced // Pass
    # "[2 ]"   #false second arg a missing      // PASS 
    # "[3  ]"   #false second arg a missing    // PASS  
    # '"[2d]"'   #false second nospace // PASS 
    # # "[[[[]]]]"  #balance but error 1) not a numebr 2) no space // PASS 
    # '"[2 3][3 4 3]3     423   423"' #true // PASS
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
    go run "$program_path" $args | tee -a "$output_file" 
    echo -e "${RED}-------END----------${REST}"
    echo ""
done

for args in "${sl_arguments[@]}"; do
    
    echo -e "${YELLOW}"running: $args
    echo ""
    echo -e "${GREEN}--------START--------${BLUE}"
    go run "$program_path" "$args" | tee -a "$output_file" 
    echo -e "${RED}-------END----------${REST}"
    echo ""
done

else 
echo "you need to install go, before running the app"
fi
