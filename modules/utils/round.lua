function round(n)
    return n % 1 >= 0.5 and math.ceil(n) or math.floor(n)
end

return round
