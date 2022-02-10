local server = {}

function Kick(playerId: PlayerID, reason: string): Response end
function Ban(playerId: PlayerID, reason: string, length: Instant): Response end
function Shutdown(): Response end

return server