local player = require("player")

function Init(gamemode)
    print("Welcome to my " .. gamemode .. " gamemode!")
end

function OnJoin(user_id)
    print("Player " .. user_id .. " has joined the match")
end

function OnLeave(user_id)
    print("Player " .. user_id .. " has left the match")
end
