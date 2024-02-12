if command -v go &> /dev/null

then
 ./run_art.sh
 ./run_server.sh
else 
echo "you need to install go, before running the app"
fi