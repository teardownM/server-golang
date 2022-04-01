#!/bin/sh

# read yaml file

TITLE=$(niet "title" ./gamemodes/config.yml)
GAMEMODE=$(niet "gamemode" ./gamemodes/config.yml)
VERSION=$(niet "version" ./gamemodes/config.yml)
DEBUG=$(niet "debug" ./gamemodes/config.yml)
MAP=$(niet "map" ./gamemodes/config.yml)

LATEST_COMMIT=$(git rev-parse HEAD)
GIT_DIFF=$(git diff --quiet HEAD) 

eval 'figlet "teardownM" | eval lolcat --animate --duration 1 --speed 55 --seed 24 --spread 24'

eval 'echo "by Alexandar Gyurov, Daniel W, Malte0621, Casinxx\n" | eval lolcat --animate --duration 1 --speed 55 --seed 24 --spread 24'

eval 'echo "################# DEV SERVER BUILD #################\n" | eval lolcat --animate --duration 1 --speed 55 --seed 24 --spread 24'

if git diff-index --quiet HEAD --; then
  echo "\e[1;33mLATEST COMMIT: $LATEST_COMMIT\n\e[0m"
else
  echo "\e[1;33m⚠️  UNCOMMITED CHANGES\e[0m"
  echo "\e[1;33mLATEST COMMIT: $LATEST_COMMIT\n\e[0m"
fi

echo "Title: $TITLE"
echo "Gamemode: $GAMEMODE" 
echo "Version: $VERSION"
echo "Debug: $DEBUG"
echo "Map: $MAP"

if [ ! -d "./modules" ]; then
  echo "\nℹ️ ./modules/ folder not found"
  eval 'mkdir modules'
  echo "✅  ./modules/ created"
fi

echo "\nℹ️  Checking if ./modules/ is empty"
if [ -z "$(ls -A ./modules)" ]; then
   echo "✅ ./modules/ is empty"
else
   eval 'rm -r modules/*'
   echo "✅  Removed old module from ./modules/"
fi

cd ./src

echo "\nℹ️  Checking if ./src/vendor/ is empty"
if [ ! -d "vendor/" ]; then
  echo "ℹ️ ./src/vendor/ not found"
  go mod vendor
  echo "✅ ./src/vendor/ created"
else
  echo "✅ ./src/vendor/ exists"
fi


echo "\nℹ️  Attempting to build Go modules from src..."
docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownM-0.0.0-unstable.so

mv ./modules/* ../modules
rm -r ./modules/
cd ..

echo "✅ Go modules built"

DOCKER_ID=$(docker ps -aqf "name=^teardownM_DEV_BUILD-nakama$")

if [ -z "$DOCKER_ID" ]
then
  echo "\n[INFO]: Could not find teardownM docker container"
  echo "\n[INFO]: Creating docker container from docker-compose.yml"
  docker-compose -p TEARDOWN_M_DEV_BUILD up
else
  echo "\n✅ Found teardownM_DEV_BUILD-nakama container with id $DOCKER_ID"
  docker restart $DOCKER_ID 

  echo "✅ Docker container restarted"
  echo "ℹ️ Connecting to container..."

  docker logs -f --tail 10 $DOCKER_ID
fi
