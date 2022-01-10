#!/bin/sh
cd ./TeardownNakamaServer

if [ $# -eq 0 ]
  then
    echo "WARNING: RUNNING THIS LOCALLY WILL HARD RESET YOUR CURRENT WORKING DIRECTORY TO MAIN!"
    echo "Add --yes to confirm"
    exit 1
fi


DIR="modules/"
if [ ! -d "$DIR" ]; then
  # Take action if $DIR exists. #
  echo "Creating module folder"
  mkdir modules
fi

git reset --hard origin/main
git pull

cd ./src
docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownNK.so

mv ./modules/* ../modules
rm -r ./modules/

cd ..

DOCKERID=$(docker ps -aqf "name=^teardownnakamaserver_nakama_1$")
docker restart $DOCKERID

echo "Successfully built and restarted Nakama"
