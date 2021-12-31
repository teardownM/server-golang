local nk = require("nakama")

local enums = require("gamemode.enums")
local player_movement = require("player.movement")

local printTable = require("utils.printTable")

function match_loop(context, dispatcher, tick, state, messages)
    -- Incoming messages
    for _, message in ipairs(messages) do
        print("\nReceived message!")

        if message.op_code == enums.OP_CODES.PLAYER_POS then
            -- A player has moved and has send a PLAYER_POS request
            local client_data = nk.json_decode(message.data)
            local client_presence = state.presences[client_data.user_id] -- find the user in presences

            print(client_data.currentX)
            --print(client_presence)

            if client_presence then
                client_presence.x = client_data.currentX
                client_presence.z = client_data.currentZ
                client_presence.y = client_data.currentY

                --printTable(client_presence)

                --print(nk.json_encode(state.presences))
                --dispatcher.broadcast_message(enums.OP_CODES.PLAYER_POS, nk.json_encode(state.presences))
            end

        -- If someone requests spawn all, send to the sender who sent it
        elseif message.op_code == enums.OP_CODES.SPAWN_ALL then
            dispatcher.broadcast_message(enums.OP_CODES.SPAWN_ALL, nk.json_encode(state.presences), {message.sender})
        end
    end

    -- TODO: Only send this if state has been updated?
    -- Every tick, send the state to all players
    dispatcher.broadcast_message(enums.OP_CODES.PLAYER_POS, nk.json_encode(state.presences))

    return state
end

return match_loop
