$Host.UI.RawUI.WindowTitle = "TeardownM - dev server";

$rebuild = $false

try {
    Import-Module powershell-yaml -ErrorAction Stop
} catch {
    write-host "[INFO]: powershell-yaml is not installed" -foreground yellow
    write-host "[INFO]: Attempting to install powershell-yaml. Please say 'YES' to all or 'NO' to cancel" -foreground yellow

    if (Read-Host -Prompt "Install powershell-yaml? (YES/NO)") {
        write-host "[INFO]: Installing powershell-yaml" -foreground yellow
        if (Install-Module powershell-yaml -Scope CurrentUser) {
            write-host "[INFO]: Successfully installed powershell-yaml" -foreground green
        } else {
            write-host "[ERROR]: Failed to install powershell-yaml" -foreground red
            exit 1
        }
    } else {
        write-host "[INFO]: Please install powershell-yaml and try again" -foreground yellow
        exit 1
    }
}

try {
    [string[]]$fileContent = Get-Content "./config.yml" -ErrorAction Stop
} catch {
    write-host "[ERROR]: Could not find ./config.yml. Please ensure the file exists." -foreground red
    exit 1
}

$content = ''
foreach ($line in $fileContent) { $content = $content + "`n" + $line }
$yaml = ConvertFrom-YAML $content

function Get-PublicIP {
    $ip_addr = "https://ipinfo.io/ip"
    $ip = (Invoke-WebRequest $ip_addr -Method Get).Content

    return $ip
}

$name = $yaml.name
$gamemode = $yaml.gamemode
$debug = $yaml.debug
$public_ip = Get-PublicIP
$map = $yaml.map

$Host.UI.RawUI.WindowTitle = "TeardownM - $name - $gamemode - $map";

if ($args[0] -eq "-rebuild") {
    $rebuild = $true
}

write-host "
  _                     _                     __  __ 
 | |                   | |                   |  \/  |
 | |_ ___  __ _ _ __ __| | _____      ___ __ | \  / |
 | __/ _ \/ _` | '__/ _` |/ _ \ \ /\ / / '_ \| |\/| |
 | |_  __/ (_| | | | (_| | (_) \ V  V /| | | | |  | |
  \__\___|\__,_|_|  \__,_|\___/ \_/\_/ |_| |_|_|  |_|

  #################### DEV BUILD ####################   

  by Alexandar Gyurov, Daniel W, Nahu and others.

  Build: 0.0.0 Unstable

  Title: $name
  Gamemode: $gamemode
  Debug: $debug
  Public IP: $public_ip
  Map: $map

  Report any bugs to: https://github.com/teardownM/server/issues

" -foreground blue 

try {
    Get-Process 'com.docker.proxy' -ErrorAction Stop > $null
} catch {
    write-host "[ERROR]: An instance of Docker could not be found. Please ensure it is running.`n" -foreground red
    Read-Host -Prompt "Press enter to exit"
    exit
}

if (-not(Test-Path -Path "modules/")) {
    write-host "[INFO]: ./modules/ folder not found" -foreground yellow
    New-Item -ItemType Directory -Force -Path "modules/"
    write-host "[SUCCESS] ./modules/ created" -foreground green
}

Set-Location .\src

if (-not(Test-Path -Path "vendor/")) {
    write-host "[INFO]: ./src/vendor/ not found" -foreground yellow
    go mod vendor
    write-host "[SUCCESS] Go vendor created" -foreground green
}

# Check if teardownM-0.0.0-unstable.so exists innside the modules folder
if (-not(Test-Path -Path "../modules/teardownM-0.0.0-unstable.so") -or $rebuild) {
    write-host "[INFO]: Attempting to build Go modules from src..." -foreground yellow
    docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:3.10.0 build -buildmode=plugin -trimpath -o ./modules/teardownM-0.0.0-unstable.so

    Move-Item -Path "./modules/*" -Destination "../modules/" -Force
    Remove-Item -Path "./modules/" -Recurse -Force

    write-host "[SUCCESS] Go modules built" -foreground green
}

if ($null -eq (docker ps -aqf "name=^teardownM_DEV_BUILD-nakama$")) {
    write-host "[INFO]: Could not find teardownM docker container" -foreground yellow
    write-host "[INFO]: Creating docker container from docker-compose.yml" -foreground yellow

    docker-compose up -d
    write-host "[SUCCESS] Docker container created" -foreground green
} else {
    write-host "[INFO]: Running teardownM_DEV_BUILD-nakama container" -foreground yellow

    docker-compose up
    # docker logs -f teardownM_DEV_BUILD-nakama
}

Set-Location ..