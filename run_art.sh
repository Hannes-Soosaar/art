if command -v go &> /dev/null
then 
go run . arg1
else 
echo "you need to install go, before running the app"
fi