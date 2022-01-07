#!/bin/sh
if [ $# -eq 0 ]
  then
    echo "Please specify the docker id to restart, find out with 'docker container ls' and copy the id of heroiclabs/nakama"
    exit 1
fi

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

docker restart $1
echo "Successfully built and restarted the Nakama server"
