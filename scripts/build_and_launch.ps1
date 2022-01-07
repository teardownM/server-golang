$dockerID = docker ps -aqf "name=^nakama-server_nakama_1$"
if ($dockerID -eq $null) {
    $dockerID2 = docker ps -aqf "name=^teardownnakamaserver-nakama-1$"

    if ($dockerID -eq $null) {
        Write-Output "No Docker ID found, please ensure Nakama and docker are running."
    }
    exit 1
}

Write-Output "Docker ID: $dockerID $dockerID2"
if (-not(Test-Path -Path "modules/")) {
    New-Item -ItemType Directory -Force -Path "modules/"
}

Set-Location .\src
docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownNK.so

Move-Item -Path "./modules/*" -Destination "../modules/" -Force
Remove-Item -Path "./modules/" -Recurse -Force

Set-Location ..

docker restart $dockerID
Write-Output "Successfully built and restarted docker container $dockerID"

docker logs -f --tail 10 $dockerID
