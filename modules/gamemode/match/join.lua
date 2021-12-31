local nk = require("nakama")

local player = require("player.settings")
local enums = require("gamemode.enums")

local printTable = require("utils.printTable")


local M = {}

function M.join_attempt(context, dispatcher, tick, state, presence, metadata)
    local acceptuser = true
    return state, acceptuser
end

function M.join(context, dispatcher, tick, state, presences)
    
    for _, presence in ipairs(presences) do
        print("[GAMEMODE] - Player has joined the game with id: ", presence.user_id)

        presence.x = 0
        presence.z = 0
        presence.y = 0

        state.presences[presence.user_id] = presence

        -- Dispatch to everyone that someone has joined
        dispatcher.broadcast_message(enums.OP_CODES.PLAYER_SPAWN, nk.json_encode(presence))
    end

    return state
end

return M
