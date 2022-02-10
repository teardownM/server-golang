local player = require("player")

function OnInit(gamemode: string)
    print("Welcome to my " .. gamemode .. " gamemode!")
end

function OnJoin(playerId: PlayerID)
    print(player.GetName(playerId) .. " has joined the match")
end

function OnLeave(playerId: PlayerID)
    print(player.GetName(playerId) .. " has left the match")
end

function OnShutdown()
    print("Server shutting down")
end

function Tick()
    
end