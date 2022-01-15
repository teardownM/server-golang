local player = require("player")

function Init(gamemode)
    print("Welcome to my " .. gamemode .. " gamemode!")
end

function OnJoin(user_id)
    print("user_id" .. user_id)
    print(player.GetHealth(user_id))
end
