PORT=4242
URL=localhost:$PORT

go build

# back ground exec
./omikuji $PORT &

JSON_RESPONSE=$(curl $URL)

echo $JSON_RESPONSE | jq "."

killall omikuji > /dev/null
