
#checks if go is installed
if command -v go &> /dev/null

then 
rm ./log/log.txt
# Array of arguments the number of arguments is the number of times the programm will be launced
arguments=(
    "-h"
    "-arg2 more non flag stuff -arg3 noflag stuff -arg4 25"
    "-non flag stuff"
)

# if the script is in the same folder, as the main go programm input "." 
program_path="./cmd/art/"

# Output file path and file name
output_file="./log/log.txt"

# Run the Go program for each set of arguments and save output to a file
for args in "${arguments[@]}"; do
    go run $program_path $args | tee -a $output_file
    # Add a sleep between runs if needed
    # sleep 1
done



else 
echo "you need to install go, before running the app"
fi