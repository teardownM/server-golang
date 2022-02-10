
local player = {}

-- Accessors
function player.GetName(playerId: PlayerID): string end
function player.GetSteamID(playerId: PlayerID): string end
function player.GetHealth(playerId: PlayerID): number end
function player.GetTransform(playerId: PlayerID): Transform end
function player.GetPosition(playerId: PlayerID): Vector3 end
function player.GetRotation(playerId: PlayerID): Quaternion end
function player.GetVelocity(playerId: PlayerID): Vector3 end
function player.GetVehicle(playerId: PlayerID): Handle | boolean end

-- Modifiers
function player.SetHealth(playerId: PlayerID, health: number): Response end
function player.SetTransform(playerId: PlayerID, transform: Transform): Response end
function player.SetPosition(playerId: PlayerID, position: Vector3): Response end
function player.SetRotation(playerId: PlayerID, transform: Quaternion): Response end
function player.SetVelocity(playerId: PlayerID, transform: Vector3): Response end
function player.SetHealth(playerId: PlayerID, transform: number): Response end

return player