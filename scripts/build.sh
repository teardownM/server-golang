#!/bin/sh

DIR="modules/"
if [ ! -d "$DIR" ]; then
  # Take action if $DIR exists. #
  echo "Creating module folder"
  mkdir modules
fi

cd ./src
docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownNK.so

mv ./modules/* ../modules -f
rm -r ./modules/

cd ..

dockerId = docker ps -aqf "name=^nakama-server_nakama_1$"
docker restart dockerId

echo "Successfully built and restarted the Nakama server"
