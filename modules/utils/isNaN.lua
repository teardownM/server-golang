--- Returns true if given value is not a number (NaN); otherwise false or nil if value is not of type string nor number.
function isNaN(value)
    if type(value) == "string" then
        value = tonumber(value)
        if value == nil then
            return nil
        end
    elseif type(value) ~= "number" then
        return nil
    end
    return value ~= value
end

return isNaN
