local nk = require("nakama")
local module = "gamemode"

local MATCH_ID = nk.match_create(module)

local function get_match_id(context)
    return nk.json_encode(MATCH_ID)
end

nk.register_rpc(get_match_id, "get_match_id")
