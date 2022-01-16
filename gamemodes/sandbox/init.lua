local player = require("player")

function Init(gamemode)
    print("Welcome to my " .. gamemode .. " gamemode!")
end

function OnJoin(user_id)
    print("user_id: " .. user_id)
    print(player.GetHealth(user_id))
    print(player.SetHealth(user_id, 50))
    print(player.GetHealth(user_id))
end
