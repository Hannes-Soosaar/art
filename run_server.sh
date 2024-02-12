#checks if go is installed
if command -v go &> /dev/null
then
rm ./log/log.txt

program_path="./cmd/server/"

output_file="./log/log.txt"

echo "starting up the server"
go run "$program_path" | tee -a "$output_file" 

else 
echo "you need to install go, before running the app"
fi