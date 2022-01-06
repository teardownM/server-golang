# TeardownNakamaServer

1. Open `docker-compose.yml`, change line 32 to match your current working directory:
`/c/Users/<username>/projects/TeardownNakamaServer:/nakama/data`

2. Launch Nakama with Docker: `docker-compose up -d`

3. Find your local Nakama Docker Container ID:
`docker-compose ls`

4. Find `heroiclabs/nakama:3.10.0` and copy it's ID

5. Build the Go files:
`./scripts/build.sh CONTAINER_ID`

6. All good to go, Nakama should have the latest module installed :) 

Follow the docs for more info:
[https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/](https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/)



