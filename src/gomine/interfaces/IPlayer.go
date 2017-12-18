package interfaces

import (
	"goraklib/server"
	"gomine/vectors"
	"gomine/entities/math"
)

type IPlayer interface {
	GetSession() *server.Session
	GetName() string
	GetDisplayName() string
	SetDisplayName(string)
	GetPermissionGroup() IPermissionGroup
	SetPermissionGroup(IPermissionGroup)
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
	GetServer() IServer
	SetViewDistance(int32)
	GetViewDistance() int32
	GetUUID() string
	GetXUID() string
	SetLanguage(string)
	GetLanguage() string
	GetClientId() int
	SetSkinId(id string)
	GetSkinId() string
	GetSkinData() []byte
	SetSkinData([]byte)
	GetCapeData() []byte
	SetCapeData([]byte)
	GetGeometryName() string
	SetGeometryName(string)
	GetGeometryData() string
	SetGeometryData(string)
	SendChunk(IChunk)
	New(IServer, *server.Session, string, string, string, int) IPlayer
	GetPing() uint64
	Move(x, y, z, pitch, yaw, headYaw float32)
	Tick()
	SendMessage(string)
	SendPacket(packet IPacket)
	PlaceInWorld(*vectors.TripleVector, *math.Rotation, ILevel, IDimension)
	GetPosition() *vectors.TripleVector
	SetPosition(*vectors.TripleVector)
	GetDimension() IDimension
	SetDimension(IDimension)
	GetLevel() ILevel
	SetLevel(ILevel)
	GetRotation() *math.Rotation
	SetRotation(*math.Rotation)
	GetMotion() *vectors.TripleVector
	SetMotion(*vectors.TripleVector)
}