function match_terminate(context, dispatcher, tick, state, grace_seconds)
    local message = "Server shutting down in " .. grace_seconds .. " seconds"
    dispatcher.broadcast_message(2, message)
    return nil
end

return match_terminate
