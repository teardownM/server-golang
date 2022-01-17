$Host.UI.RawUI.WindowTitle = "TeardownM - dev server";

try {
    Import-Module powershell-yaml -ErrorAction Stop
}
catch {
    write-host "[INFO]: powershell-yaml is not installed" -foreground yellow
    write-host "[INFO]: Attempting to install powershell-yaml. Please say 'YES' to all" -foreground yellow
    Install-Module powershell-yaml -Scope CurrentUser
    write-host "[SUCCESS]: Installed 'powershell-yaml'" -foreground green
}

try {
    [string[]]$fileContent = Get-Content "./gamemodes/config.yml" -ErrorAction Stop
}
catch {
    write-host "[ERROR]: Could not find ./gamemodes/config.yml. Please ensure the file exists." -foreground red
}

$content = ''
foreach ($line in $fileContent) { $content = $content + "`n" + $line }
$yaml = ConvertFrom-YAML $content

$title = $yaml.title
$gamemode = $yaml.gamemode
$version = $yaml.version
$debug = $yaml.debug
$public_ip = $yaml.public_ip
$map = $yaml.map

$Host.UI.RawUI.WindowTitle = "TeardownM - $title - $gamemode - $map - v$version";

write-host "
  _                     _                     __  __ 
 | |                   | |                   |  \/  |
 | |_ ___  __ _ _ __ __| | _____      ___ __ | \  / |
 | __/ _ \/ _` | '__/ _` |/ _ \ \ /\ / / '_ \| |\/| |
 | |_  __/ (_| | | | (_| | (_) \ V  V /| | | | |  | |
  \__\___|\__,_|_|  \__,_|\___/ \_/\_/ |_| |_|_|  |_|

  #################### DEV BUILD ####################   

  by Alexandar Gyurov, Daniel W, Malte0621, Casin

  Build: 0.0.0 unstable

  Title: $title
  Gamemode: $gamemode
  Version: $version
  Debug: $debug
  Public IP: $public_ip
  Map: $map

  Report any bugs to: https://github.com/teardownM/server/issues

" -foreground blue 

try {
    Get-Process 'com.docker.proxy' -ErrorAction Stop > $null
    write-host "[SUCCESS] Found docker" -foreground green
}
catch {
    write-host "[ERROR]: An instance of Docker could not be found. Please ensure it is running.`n" -foreground red
    Read-Host -Prompt "Press enter to exit"
    exit
}

if (-not(Test-Path -Path "modules/")) {
    write-host "[INFO]: ./modules/ folder not found" -foreground yellow
    New-Item -ItemType Directory -Force -Path "modules/"
    write-host "[SUCCESS] ./modules/ created" -foreground green
}


Remove-Item ./modules/* -Recurse -Force
write-host "[SUCCESS] Removed old module from ./modules/" -foreground green

Set-Location .\src

if (-not(Test-Path -Path "vendor/")) {
    write-host "[INFO]: ./src/vendor/ not found" -foreground yellow
    go mod vendor
    write-host "[SUCCESS] Go vendor created" -foreground green
}

write-host "[INFO]: Attempting to build Go modules from src..." -foreground yellow
docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownM-0.0.0-unstable.so

Move-Item -Path "./modules/*" -Destination "../modules/" -Force
Remove-Item -Path "./modules/" -Recurse -Force
Set-Location ..

write-host "[SUCCESS] Go modules built" -foreground green

if ($null -eq (docker ps -aqf "name=^teardownM_DEV_BUILD-nakama$")) {
    write-host "[INFO]: Could not find teardownM docker container" -foreground yellow
    write-host "[INFO]: Creating docker container from docker-compose.yml" -foreground yellow

    docker-compose -p TEARDOWN_M_DEV_BUILD up
    Read-Host -Prompt "Press enter to exit"
}
else {
    write-host "[SUCCESS] Found teardownM_DEV_BUILD-nakama container" -foreground green
    docker restart (docker ps -aqf "name=^teardownM_DEV_BUILD-nakama$") | Out-Null
    write-host "[SUCCESS] Docker container restarted" -foreground green
    write-host "[INFO] Connecting to container..." -foreground yellow

    docker logs -f --tail 10 (docker ps -aqf "name=^teardownM_DEV_BUILD-nakama$" )
}
