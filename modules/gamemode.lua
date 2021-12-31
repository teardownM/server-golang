local M = {}

M.match_init = require("gamemode.match.init")

M.match_join = require("gamemode.match.join").join
M.match_join_attempt = require("gamemode.match.join").join_attempt
M.match_leave = require("gamemode.match.leave")
M.match_loop = require("gamemode.match.loop")
M.match_terminate = require("gamemode.match.terminate")

return M
