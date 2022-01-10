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

DOCKERID=$(docker ps -aqf "name=^teardownnakamaserver_nakama_1$")
docker restart $DOCKERID

echo "Successfully built and restarted the Nakama server"
