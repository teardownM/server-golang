# TeardownNakamaServer

1. Open `docker-compose.yml`, change line 32 to match your current working directory:
`/c/Users/<username>/projects/TeardownNakamaServer:/nakama/data`

2. To build the Go files:
`docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownNK.so`

3. Move the modules folder in parent directory (outside of src) (haven't figured out how to automatically do this)

4. Run `docker compose up`

5. Use docker desktop to restart just the Nakama server to refresh the server files on every build

Follow the docs for more info:
[https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/](https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/)



