local player = require("player")

function OnInit(gamemode)
   print("Welcome to my " .. gamemode .. " gamemode!")
end

function OnJoin(playerId)
   print(player.GetName(playerId) .. " has joined the match")
end

function OnLeave(playerId)
   print(player.GetName(playerId) .. " has left the match")
end

function OnShutdown()
   print("Server shutting down")
end

function Tick()

end
