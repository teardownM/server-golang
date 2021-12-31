function match_leave(context, dispatcher, tick, state, presences)
    for _, presence in ipairs(presences) do
        state.presences[presence.user_id] = nil
    end
    return state
end

return match_leave
