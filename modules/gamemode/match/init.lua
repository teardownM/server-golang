function match_init(context, setupstate)
    local tickrate = 28
    local label = "main"
    local gamestate = {
        presences = {}
    }

    return gamestate, tickrate, label
end

return match_init
