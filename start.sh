containerID=$(docker ps -aqf "name=^nakama-server_nakama_1$")

./scripts/build.sh $containerID

docker logs -f --tail 10 $containerID
