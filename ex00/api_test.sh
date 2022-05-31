PORT=4242
URL=localhost:$PORT


# install jq
which jq > /dev/null
if [ $? == 1 ]; then
	brew install jq
fi

# build
go build

# start server
./omikuji $PORT &

# send request
JSON_RESPONSE=$(curl $URL)
echo $JSON_RESPONSE | jq "."

# end server
killall omikuji > /dev/null
