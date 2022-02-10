global type PlayerID = string
global type Handle = number
global type Instant = string -- an ISO date

global type Vector3 = {number, number, number}
global type Quaternion = {number, number, number, number}

global record Transform
    pos: Vector3
    rot: Quaternion
end

global record Response
   success: boolean
   err: string | nil
end

--===============================--
--======= CALLBACKS TYPES =======--
--===============================--

global OnInit: function(gamemode: string): nil 
global OnJoin: function(playerId: string): nil 
global OnLeave: function(playerId: string): nil 
global OnShutdown: function(): nil 

global Tick: function(): nil 

--============================--
--======= PLAYER TYPES =======--
--============================--

global record player
    -- Accessors
    GetName: function(playerId: PlayerID): string
    GetSteamID: function(playerId: PlayerID): string
    GetHealth: function(playerId: PlayerID): number
    GetTransform: function(playerId: PlayerID): Transform
    GetPosition: function(playerId: PlayerID): Vector3
    GetRotation: function(playerId: PlayerID): Quaternion
    GetVelocity: function(playerId: PlayerID): Vector3
    GetVehicle: function(playerId: PlayerID): Handle | boolean

    -- Modifiers
    SetHealth: function(playerId: PlayerID, health: number): Response
    SetTransform: function(playerId: PlayerID, transform: Transform): Response
    SetPosition: function(playerId: PlayerID, position: Vector3): Response
    SetRotation: function(playerId: PlayerID, transform: Quaternion): Response
    SetVelocity: function(playerId: PlayerID, transform: Vector3): Response
    SetHealth: function(playerId: PlayerID, transform: number): Response
end

--============================--
--======= SERVER TYPES =======--
--============================--

global record server
    Kick: function(playerId: PlayerID, reason: string): Response
    Ban: function(playerId: PlayerID, reason: string, length: Instant): Response
    Shutdown: function(): Response
end
