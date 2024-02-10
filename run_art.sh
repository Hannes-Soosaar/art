
#checks if go is installed
if command -v go &> /dev/null

then 
rm ./log/log.txt
# Array of arguments the number of arguments is the number of times the programm will be launced
arguments=(
    # "-h" #PASS
    # "-e" # PASS 
    # "-m assets/input/encoded/cats.encoded.txt" #PASS
    # "-m assets/input/encoded/kood.encoded.txt" #PASS
    # "-m assets/input/encoded/lion.encoded.txt" #PASS
    # "-m assets/input/encoded/plane.encoded.txt" #PASS
    # "-e --15454111115445454AAAAAAAAAAAABBBBBBBBBBBCCCCDE" #PASS
    # "-m -e filetoencode.txt"
    # "-m -e filetoencode2.txt"
    # "-m -2d filetoencode3.txt"
#     -e "[45 ][334]"
#     "-m -e filetoencode.txt " #FAIL
#     "random text to see if it will triger help"  #FAIL // Check for max amount of chars
#     -e "[45 ][334]"  # test encode // FAIL  // Bug it will still do the same validation for encoding ?
)

sl_arguments=(
    # "[5 #][5 -_]-[5 #]" # PASS
    # "ABC[10 D]EFG" # PASS
    # "asd][5 s][gf]" #false unbalance brackets //PASS
    # '"[a df]fd23545[afd3253]"'   #false first arg not a number 
    # "[15 #]       [5 -_] - f d [5 #]"   #PASS
    # "[19 D]"   #PASS
    #  '"[2    d"'   #false unbalanced #PASS
    # "[2 ]"   #false second arg a missing      // PASS 
    # "[3  ]"   #PASS 
    # '"[2d]"'   #false second nospace // PASS 
    # "[[[[]]]]"  #balance but error 1) not a numebr 2) no space // PASS 
    # '"[2 3][3 4 3]3     423   423"' #true // PASS
)

vect_arguments=(
    '-e "45dddddddddddddddddd fdsaffasdffffffff334"'
    '-e "abcdefghijklmn op galskdfj sdfa"'  # test encode // FAIL  // Bug it will still do the same validation for encoding 
    '-e "[][[[[[]][][[]]]]]"'  # test encode // FAIL  // Bug it will still do the same validation for encoding 
    '-e "abcdefghijklmn op dddfdfff  222222222222221 ***-*-*-*-**-*-************galskdfj sdfa"'  # test encode // FAIL  // Bug it will still do the same validation for encoding 
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


#Run a programm so taht there are not spaces in any of the args 
for args_set in "${arguments[@]}"; do
    args=($arg_set)
 
    echo -e "${YELLOW}" running: $args_set
    # echo -e "${YELLOW}"running: "${arguments[@]}"
    echo ""
    echo -e "${GREEN}--------START--------${BLUE}"
    go run "$program_path" $args_set | tee -a "$output_file" 
    echo -e "${RED}-------END----------${REST}"
    echo ""
done
# runs the programm such that the second argument can have spaces 'arg1 "arg 2 with spaces "'
for arg_set in "${vect_arguments[@]}"; do
    eval "args=($arg_set)"
    arg1=${args[0]}
    arg2=${args[1]}
    echo -e "${YELLOW}"running: "$arg1 \"$arg2\""
    echo ""
    echo -e "${GREEN}--------START--------${BLUE}"
    go run "$program_path" "${args[@]}" | tee -a "$output_file"
    echo -e "${RED}-------END----------${REST}"
    echo ""
done
# Run the programm such that it has a single argument that can contain spaces
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
