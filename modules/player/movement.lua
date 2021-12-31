function move_player(presence, client_data)
    -- TODO: Perhaps I can optimise this function?

    local accelerationX = presence.forceX / presence.mass
    local accelerationY = presence.forceY / presence.mass

    local MAX_FORCE = presence.mass * 10
    -- Check not going faster than max force
    if presence.forceX >= MAX_FORCE then
        presence.forceX = MAX_FORCE
    elseif presence.forceX <= -MAX_FORCE then
        presence.forceX = -MAX_FORCE
    end

    if presence.forceY >= MAX_FORCE then
        presence.forceY = MAX_FORCE
    elseif presence.forceY <= -MAX_FORCE then
        presence.forceY = -MAX_FORCE
    end

    -- Apply force
    if client_data.moveX == 1 then
        presence.forceX = presence.forceX + 1
        presence.x = presence.x + accelerationX
    end

    if client_data.moveX == -1 then
        presence.forceX = presence.forceX - 1
        presence.x = presence.x + accelerationX
    end

    if client_data.moveY == 1 then
        presence.forceY = presence.forceY + 1
        presence.y = presence.y + accelerationY
    end

    if client_data.moveY == -1 then
        presence.forceY = presence.forceY - 1
        presence.y = presence.y + accelerationY
    end

    -- Slow down the player
    if client_data.moveX == 0 then
        if presence.forceX > 0 then
            presence.forceX = presence.forceX - 1
            presence.x = presence.x + accelerationX
        elseif presence.forceX < 0 then
            presence.forceX = presence.forceX + 1
            presence.x = presence.x + accelerationX
        end
    end

    if client_data.moveY == 0 then
        if presence.forceY > 0 then
            presence.forceY = presence.forceY - 1
            presence.y = presence.y + accelerationY
        elseif presence.forceY < 0 then
            presence.forceY = presence.forceY + 1
            presence.y = presence.y + accelerationY
        end
    end
end

return move_player
