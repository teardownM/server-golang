# TeardownNakamaServer

Follow the docs to install Nakama with Docker
[https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/](https://heroiclabs.com/docs/nakama/getting-started/docker-quickstart/)

Make sure you change the volume to your directory eg:

`/c/Users/<username>/projects/TeardownNakamaServer:/nakama/data` on line 30 in your `docker-compose.yml` file

Run `docker compose up`

Nakama should be running now. Should be able to run the Sledge Mod.

Take a look at the lua files, the main loop gets run in `loop.lua`. Restart the server (just the nakama server not the databases) to refresh any changes you make, you can use docker desktop to do it easily.
