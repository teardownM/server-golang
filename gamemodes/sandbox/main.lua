local t = require("gamemodes.sandbox.test")

function OnInitialize()
    LogGeneral("Sandbox gamemode loaded!")
end

function OnConnected(user_id)
    LogGeneral("Player '%s' connected!", user_id)
end

function OnDisconnected(user_id)
    LogGeneral("Player '%s' has left the match", user_id)
end

function OnTick(tick)
    -- LogGeneral("Tick %s", tick)
end