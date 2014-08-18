// Code generated by protoc-gen-gogo.
// source: demo.proto
// DO NOT EDIT!

package dota

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import google_protobuf "github.com/dotabuff/sange/dota/google/protobuf/descriptor.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type EDemoCommands int32

const (
	EDemoCommands_DEM_Error               EDemoCommands = -1
	EDemoCommands_DEM_Stop                EDemoCommands = 0
	EDemoCommands_DEM_FileHeader          EDemoCommands = 1
	EDemoCommands_DEM_FileInfo            EDemoCommands = 2
	EDemoCommands_DEM_SyncTick            EDemoCommands = 3
	EDemoCommands_DEM_SendTables          EDemoCommands = 4
	EDemoCommands_DEM_ClassInfo           EDemoCommands = 5
	EDemoCommands_DEM_StringTables        EDemoCommands = 6
	EDemoCommands_DEM_Packet              EDemoCommands = 7
	EDemoCommands_DEM_SignonPacket        EDemoCommands = 8
	EDemoCommands_DEM_ConsoleCmd          EDemoCommands = 9
	EDemoCommands_DEM_CustomData          EDemoCommands = 10
	EDemoCommands_DEM_CustomDataCallbacks EDemoCommands = 11
	EDemoCommands_DEM_UserCmd             EDemoCommands = 12
	EDemoCommands_DEM_FullPacket          EDemoCommands = 13
	EDemoCommands_DEM_SaveGame            EDemoCommands = 14
	EDemoCommands_DEM_SpawnGroups         EDemoCommands = 15
	EDemoCommands_DEM_Max                 EDemoCommands = 16
	EDemoCommands_DEM_IsCompressed        EDemoCommands = 64
)

var EDemoCommands_name = map[int32]string{
	-1: "DEM_Error",
	0:  "DEM_Stop",
	1:  "DEM_FileHeader",
	2:  "DEM_FileInfo",
	3:  "DEM_SyncTick",
	4:  "DEM_SendTables",
	5:  "DEM_ClassInfo",
	6:  "DEM_StringTables",
	7:  "DEM_Packet",
	8:  "DEM_SignonPacket",
	9:  "DEM_ConsoleCmd",
	10: "DEM_CustomData",
	11: "DEM_CustomDataCallbacks",
	12: "DEM_UserCmd",
	13: "DEM_FullPacket",
	14: "DEM_SaveGame",
	15: "DEM_SpawnGroups",
	16: "DEM_Max",
	64: "DEM_IsCompressed",
}
var EDemoCommands_value = map[string]int32{
	"DEM_Error":               -1,
	"DEM_Stop":                0,
	"DEM_FileHeader":          1,
	"DEM_FileInfo":            2,
	"DEM_SyncTick":            3,
	"DEM_SendTables":          4,
	"DEM_ClassInfo":           5,
	"DEM_StringTables":        6,
	"DEM_Packet":              7,
	"DEM_SignonPacket":        8,
	"DEM_ConsoleCmd":          9,
	"DEM_CustomData":          10,
	"DEM_CustomDataCallbacks": 11,
	"DEM_UserCmd":             12,
	"DEM_FullPacket":          13,
	"DEM_SaveGame":            14,
	"DEM_SpawnGroups":         15,
	"DEM_Max":                 16,
	"DEM_IsCompressed":        64,
}

func (x EDemoCommands) Enum() *EDemoCommands {
	p := new(EDemoCommands)
	*p = x
	return p
}
func (x EDemoCommands) String() string {
	return proto.EnumName(EDemoCommands_name, int32(x))
}
func (x *EDemoCommands) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDemoCommands_value, data, "EDemoCommands")
	if err != nil {
		return err
	}
	*x = EDemoCommands(value)
	return nil
}

type CDemoFileHeader struct {
	DemoFileStamp            *string `protobuf:"bytes,1,req,name=demo_file_stamp" json:"demo_file_stamp,omitempty"`
	NetworkProtocol          *int32  `protobuf:"varint,2,opt,name=network_protocol" json:"network_protocol,omitempty"`
	ServerName               *string `protobuf:"bytes,3,opt,name=server_name" json:"server_name,omitempty"`
	ClientName               *string `protobuf:"bytes,4,opt,name=client_name" json:"client_name,omitempty"`
	MapName                  *string `protobuf:"bytes,5,opt,name=map_name" json:"map_name,omitempty"`
	GameDirectory            *string `protobuf:"bytes,6,opt,name=game_directory" json:"game_directory,omitempty"`
	FullpacketsVersion       *int32  `protobuf:"varint,7,opt,name=fullpackets_version" json:"fullpackets_version,omitempty"`
	AllowClientsideEntities  *bool   `protobuf:"varint,8,opt,name=allow_clientside_entities" json:"allow_clientside_entities,omitempty"`
	AllowClientsideParticles *bool   `protobuf:"varint,9,opt,name=allow_clientside_particles" json:"allow_clientside_particles,omitempty"`
	XXX_unrecognized         []byte  `json:"-"`
}

func (m *CDemoFileHeader) Reset()         { *m = CDemoFileHeader{} }
func (m *CDemoFileHeader) String() string { return proto.CompactTextString(m) }
func (*CDemoFileHeader) ProtoMessage()    {}

func (m *CDemoFileHeader) GetDemoFileStamp() string {
	if m != nil && m.DemoFileStamp != nil {
		return *m.DemoFileStamp
	}
	return ""
}

func (m *CDemoFileHeader) GetNetworkProtocol() int32 {
	if m != nil && m.NetworkProtocol != nil {
		return *m.NetworkProtocol
	}
	return 0
}

func (m *CDemoFileHeader) GetServerName() string {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return ""
}

func (m *CDemoFileHeader) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

func (m *CDemoFileHeader) GetMapName() string {
	if m != nil && m.MapName != nil {
		return *m.MapName
	}
	return ""
}

func (m *CDemoFileHeader) GetGameDirectory() string {
	if m != nil && m.GameDirectory != nil {
		return *m.GameDirectory
	}
	return ""
}

func (m *CDemoFileHeader) GetFullpacketsVersion() int32 {
	if m != nil && m.FullpacketsVersion != nil {
		return *m.FullpacketsVersion
	}
	return 0
}

func (m *CDemoFileHeader) GetAllowClientsideEntities() bool {
	if m != nil && m.AllowClientsideEntities != nil {
		return *m.AllowClientsideEntities
	}
	return false
}

func (m *CDemoFileHeader) GetAllowClientsideParticles() bool {
	if m != nil && m.AllowClientsideParticles != nil {
		return *m.AllowClientsideParticles
	}
	return false
}

type CGameInfo struct {
	Dota             *CGameInfo_CDotaGameInfo `protobuf:"bytes,4,opt,name=dota" json:"dota,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *CGameInfo) Reset()         { *m = CGameInfo{} }
func (m *CGameInfo) String() string { return proto.CompactTextString(m) }
func (*CGameInfo) ProtoMessage()    {}

func (m *CGameInfo) GetDota() *CGameInfo_CDotaGameInfo {
	if m != nil {
		return m.Dota
	}
	return nil
}

type CGameInfo_CDotaGameInfo struct {
	MatchId          *uint32                                     `protobuf:"varint,1,opt,name=match_id" json:"match_id,omitempty"`
	GameMode         *int32                                      `protobuf:"varint,2,opt,name=game_mode" json:"game_mode,omitempty"`
	GameWinner       *int32                                      `protobuf:"varint,3,opt,name=game_winner" json:"game_winner,omitempty"`
	PlayerInfo       []*CGameInfo_CDotaGameInfo_CPlayerInfo      `protobuf:"bytes,4,rep,name=player_info" json:"player_info,omitempty"`
	Leagueid         *uint32                                     `protobuf:"varint,5,opt,name=leagueid" json:"leagueid,omitempty"`
	PicksBans        []*CGameInfo_CDotaGameInfo_CHeroSelectEvent `protobuf:"bytes,6,rep,name=picks_bans" json:"picks_bans,omitempty"`
	RadiantTeamId    *uint32                                     `protobuf:"varint,7,opt,name=radiant_team_id" json:"radiant_team_id,omitempty"`
	DireTeamId       *uint32                                     `protobuf:"varint,8,opt,name=dire_team_id" json:"dire_team_id,omitempty"`
	RadiantTeamTag   *string                                     `protobuf:"bytes,9,opt,name=radiant_team_tag" json:"radiant_team_tag,omitempty"`
	DireTeamTag      *string                                     `protobuf:"bytes,10,opt,name=dire_team_tag" json:"dire_team_tag,omitempty"`
	EndTime          *uint32                                     `protobuf:"varint,11,opt,name=end_time" json:"end_time,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *CGameInfo_CDotaGameInfo) Reset()         { *m = CGameInfo_CDotaGameInfo{} }
func (m *CGameInfo_CDotaGameInfo) String() string { return proto.CompactTextString(m) }
func (*CGameInfo_CDotaGameInfo) ProtoMessage()    {}

func (m *CGameInfo_CDotaGameInfo) GetMatchId() uint32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetGameMode() int32 {
	if m != nil && m.GameMode != nil {
		return *m.GameMode
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetGameWinner() int32 {
	if m != nil && m.GameWinner != nil {
		return *m.GameWinner
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetPlayerInfo() []*CGameInfo_CDotaGameInfo_CPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *CGameInfo_CDotaGameInfo) GetLeagueid() uint32 {
	if m != nil && m.Leagueid != nil {
		return *m.Leagueid
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetPicksBans() []*CGameInfo_CDotaGameInfo_CHeroSelectEvent {
	if m != nil {
		return m.PicksBans
	}
	return nil
}

func (m *CGameInfo_CDotaGameInfo) GetRadiantTeamId() uint32 {
	if m != nil && m.RadiantTeamId != nil {
		return *m.RadiantTeamId
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetDireTeamId() uint32 {
	if m != nil && m.DireTeamId != nil {
		return *m.DireTeamId
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo) GetRadiantTeamTag() string {
	if m != nil && m.RadiantTeamTag != nil {
		return *m.RadiantTeamTag
	}
	return ""
}

func (m *CGameInfo_CDotaGameInfo) GetDireTeamTag() string {
	if m != nil && m.DireTeamTag != nil {
		return *m.DireTeamTag
	}
	return ""
}

func (m *CGameInfo_CDotaGameInfo) GetEndTime() uint32 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type CGameInfo_CDotaGameInfo_CPlayerInfo struct {
	HeroName         *string `protobuf:"bytes,1,opt,name=hero_name" json:"hero_name,omitempty"`
	PlayerName       *string `protobuf:"bytes,2,opt,name=player_name" json:"player_name,omitempty"`
	IsFakeClient     *bool   `protobuf:"varint,3,opt,name=is_fake_client" json:"is_fake_client,omitempty"`
	Steamid          *uint64 `protobuf:"varint,4,opt,name=steamid" json:"steamid,omitempty"`
	GameTeam         *int32  `protobuf:"varint,5,opt,name=game_team" json:"game_team,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) Reset()         { *m = CGameInfo_CDotaGameInfo_CPlayerInfo{} }
func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) String() string { return proto.CompactTextString(m) }
func (*CGameInfo_CDotaGameInfo_CPlayerInfo) ProtoMessage()    {}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) GetHeroName() string {
	if m != nil && m.HeroName != nil {
		return *m.HeroName
	}
	return ""
}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) GetIsFakeClient() bool {
	if m != nil && m.IsFakeClient != nil {
		return *m.IsFakeClient
	}
	return false
}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo_CPlayerInfo) GetGameTeam() int32 {
	if m != nil && m.GameTeam != nil {
		return *m.GameTeam
	}
	return 0
}

type CGameInfo_CDotaGameInfo_CHeroSelectEvent struct {
	IsPick           *bool   `protobuf:"varint,1,opt,name=is_pick" json:"is_pick,omitempty"`
	Team             *uint32 `protobuf:"varint,2,opt,name=team" json:"team,omitempty"`
	HeroId           *uint32 `protobuf:"varint,3,opt,name=hero_id" json:"hero_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGameInfo_CDotaGameInfo_CHeroSelectEvent) Reset() {
	*m = CGameInfo_CDotaGameInfo_CHeroSelectEvent{}
}
func (m *CGameInfo_CDotaGameInfo_CHeroSelectEvent) String() string { return proto.CompactTextString(m) }
func (*CGameInfo_CDotaGameInfo_CHeroSelectEvent) ProtoMessage()    {}

func (m *CGameInfo_CDotaGameInfo_CHeroSelectEvent) GetIsPick() bool {
	if m != nil && m.IsPick != nil {
		return *m.IsPick
	}
	return false
}

func (m *CGameInfo_CDotaGameInfo_CHeroSelectEvent) GetTeam() uint32 {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return 0
}

func (m *CGameInfo_CDotaGameInfo_CHeroSelectEvent) GetHeroId() uint32 {
	if m != nil && m.HeroId != nil {
		return *m.HeroId
	}
	return 0
}

type CDemoFileInfo struct {
	PlaybackTime     *float32   `protobuf:"fixed32,1,opt,name=playback_time" json:"playback_time,omitempty"`
	PlaybackTicks    *int32     `protobuf:"varint,2,opt,name=playback_ticks" json:"playback_ticks,omitempty"`
	PlaybackFrames   *int32     `protobuf:"varint,3,opt,name=playback_frames" json:"playback_frames,omitempty"`
	GameInfo         *CGameInfo `protobuf:"bytes,4,opt,name=game_info" json:"game_info,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *CDemoFileInfo) Reset()         { *m = CDemoFileInfo{} }
func (m *CDemoFileInfo) String() string { return proto.CompactTextString(m) }
func (*CDemoFileInfo) ProtoMessage()    {}

func (m *CDemoFileInfo) GetPlaybackTime() float32 {
	if m != nil && m.PlaybackTime != nil {
		return *m.PlaybackTime
	}
	return 0
}

func (m *CDemoFileInfo) GetPlaybackTicks() int32 {
	if m != nil && m.PlaybackTicks != nil {
		return *m.PlaybackTicks
	}
	return 0
}

func (m *CDemoFileInfo) GetPlaybackFrames() int32 {
	if m != nil && m.PlaybackFrames != nil {
		return *m.PlaybackFrames
	}
	return 0
}

func (m *CDemoFileInfo) GetGameInfo() *CGameInfo {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

type CDemoPacket struct {
	SequenceIn       *int32 `protobuf:"varint,1,opt,name=sequence_in" json:"sequence_in,omitempty"`
	SequenceOutAck   *int32 `protobuf:"varint,2,opt,name=sequence_out_ack" json:"sequence_out_ack,omitempty"`
	Data             []byte `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoPacket) Reset()         { *m = CDemoPacket{} }
func (m *CDemoPacket) String() string { return proto.CompactTextString(m) }
func (*CDemoPacket) ProtoMessage()    {}

func (m *CDemoPacket) GetSequenceIn() int32 {
	if m != nil && m.SequenceIn != nil {
		return *m.SequenceIn
	}
	return 0
}

func (m *CDemoPacket) GetSequenceOutAck() int32 {
	if m != nil && m.SequenceOutAck != nil {
		return *m.SequenceOutAck
	}
	return 0
}

func (m *CDemoPacket) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDemoFullPacket struct {
	StringTable      *CDemoStringTables `protobuf:"bytes,1,opt,name=string_table" json:"string_table,omitempty"`
	Packet           *CDemoPacket       `protobuf:"bytes,2,opt,name=packet" json:"packet,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CDemoFullPacket) Reset()         { *m = CDemoFullPacket{} }
func (m *CDemoFullPacket) String() string { return proto.CompactTextString(m) }
func (*CDemoFullPacket) ProtoMessage()    {}

func (m *CDemoFullPacket) GetStringTable() *CDemoStringTables {
	if m != nil {
		return m.StringTable
	}
	return nil
}

func (m *CDemoFullPacket) GetPacket() *CDemoPacket {
	if m != nil {
		return m.Packet
	}
	return nil
}

type CDemoSaveGame struct {
	Data             []byte  `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	SteamId          *uint64 `protobuf:"fixed64,2,opt,name=steam_id" json:"steam_id,omitempty"`
	Signature        *uint64 `protobuf:"fixed64,3,opt,name=signature" json:"signature,omitempty"`
	Version          *int32  `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDemoSaveGame) Reset()         { *m = CDemoSaveGame{} }
func (m *CDemoSaveGame) String() string { return proto.CompactTextString(m) }
func (*CDemoSaveGame) ProtoMessage()    {}

func (m *CDemoSaveGame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CDemoSaveGame) GetSteamId() uint64 {
	if m != nil && m.SteamId != nil {
		return *m.SteamId
	}
	return 0
}

func (m *CDemoSaveGame) GetSignature() uint64 {
	if m != nil && m.Signature != nil {
		return *m.Signature
	}
	return 0
}

func (m *CDemoSaveGame) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type CDemoSyncTick struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoSyncTick) Reset()         { *m = CDemoSyncTick{} }
func (m *CDemoSyncTick) String() string { return proto.CompactTextString(m) }
func (*CDemoSyncTick) ProtoMessage()    {}

type CDemoConsoleCmd struct {
	Cmdstring        *string `protobuf:"bytes,1,opt,name=cmdstring" json:"cmdstring,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDemoConsoleCmd) Reset()         { *m = CDemoConsoleCmd{} }
func (m *CDemoConsoleCmd) String() string { return proto.CompactTextString(m) }
func (*CDemoConsoleCmd) ProtoMessage()    {}

func (m *CDemoConsoleCmd) GetCmdstring() string {
	if m != nil && m.Cmdstring != nil {
		return *m.Cmdstring
	}
	return ""
}

type CDemoSendTables struct {
	Data             []byte `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoSendTables) Reset()         { *m = CDemoSendTables{} }
func (m *CDemoSendTables) String() string { return proto.CompactTextString(m) }
func (*CDemoSendTables) ProtoMessage()    {}

func (m *CDemoSendTables) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDemoClassInfo struct {
	Classes          []*CDemoClassInfoClassT `protobuf:"bytes,1,rep,name=classes" json:"classes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *CDemoClassInfo) Reset()         { *m = CDemoClassInfo{} }
func (m *CDemoClassInfo) String() string { return proto.CompactTextString(m) }
func (*CDemoClassInfo) ProtoMessage()    {}

func (m *CDemoClassInfo) GetClasses() []*CDemoClassInfoClassT {
	if m != nil {
		return m.Classes
	}
	return nil
}

type CDemoClassInfoClassT struct {
	ClassId          *int32  `protobuf:"varint,1,opt,name=class_id" json:"class_id,omitempty"`
	NetworkName      *string `protobuf:"bytes,2,opt,name=network_name" json:"network_name,omitempty"`
	TableName        *string `protobuf:"bytes,3,opt,name=table_name" json:"table_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDemoClassInfoClassT) Reset()         { *m = CDemoClassInfoClassT{} }
func (m *CDemoClassInfoClassT) String() string { return proto.CompactTextString(m) }
func (*CDemoClassInfoClassT) ProtoMessage()    {}

func (m *CDemoClassInfoClassT) GetClassId() int32 {
	if m != nil && m.ClassId != nil {
		return *m.ClassId
	}
	return 0
}

func (m *CDemoClassInfoClassT) GetNetworkName() string {
	if m != nil && m.NetworkName != nil {
		return *m.NetworkName
	}
	return ""
}

func (m *CDemoClassInfoClassT) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

type CDemoCustomData struct {
	CallbackIndex    *int32 `protobuf:"varint,1,opt,name=callback_index" json:"callback_index,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoCustomData) Reset()         { *m = CDemoCustomData{} }
func (m *CDemoCustomData) String() string { return proto.CompactTextString(m) }
func (*CDemoCustomData) ProtoMessage()    {}

func (m *CDemoCustomData) GetCallbackIndex() int32 {
	if m != nil && m.CallbackIndex != nil {
		return *m.CallbackIndex
	}
	return 0
}

func (m *CDemoCustomData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDemoCustomDataCallbacks struct {
	SaveId           []string `protobuf:"bytes,1,rep,name=save_id" json:"save_id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDemoCustomDataCallbacks) Reset()         { *m = CDemoCustomDataCallbacks{} }
func (m *CDemoCustomDataCallbacks) String() string { return proto.CompactTextString(m) }
func (*CDemoCustomDataCallbacks) ProtoMessage()    {}

func (m *CDemoCustomDataCallbacks) GetSaveId() []string {
	if m != nil {
		return m.SaveId
	}
	return nil
}

type CDemoStringTables struct {
	Tables           []*CDemoStringTablesTableT `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CDemoStringTables) Reset()         { *m = CDemoStringTables{} }
func (m *CDemoStringTables) String() string { return proto.CompactTextString(m) }
func (*CDemoStringTables) ProtoMessage()    {}

func (m *CDemoStringTables) GetTables() []*CDemoStringTablesTableT {
	if m != nil {
		return m.Tables
	}
	return nil
}

type CDemoStringTablesItemsT struct {
	Str              *string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDemoStringTablesItemsT) Reset()         { *m = CDemoStringTablesItemsT{} }
func (m *CDemoStringTablesItemsT) String() string { return proto.CompactTextString(m) }
func (*CDemoStringTablesItemsT) ProtoMessage()    {}

func (m *CDemoStringTablesItemsT) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func (m *CDemoStringTablesItemsT) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDemoStringTablesTableT struct {
	TableName        *string                    `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	Items            []*CDemoStringTablesItemsT `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	ItemsClientside  []*CDemoStringTablesItemsT `protobuf:"bytes,3,rep,name=items_clientside" json:"items_clientside,omitempty"`
	TableFlags       *int32                     `protobuf:"varint,4,opt,name=table_flags" json:"table_flags,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CDemoStringTablesTableT) Reset()         { *m = CDemoStringTablesTableT{} }
func (m *CDemoStringTablesTableT) String() string { return proto.CompactTextString(m) }
func (*CDemoStringTablesTableT) ProtoMessage()    {}

func (m *CDemoStringTablesTableT) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

func (m *CDemoStringTablesTableT) GetItems() []*CDemoStringTablesItemsT {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CDemoStringTablesTableT) GetItemsClientside() []*CDemoStringTablesItemsT {
	if m != nil {
		return m.ItemsClientside
	}
	return nil
}

func (m *CDemoStringTablesTableT) GetTableFlags() int32 {
	if m != nil && m.TableFlags != nil {
		return *m.TableFlags
	}
	return 0
}

type CDemoStop struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoStop) Reset()         { *m = CDemoStop{} }
func (m *CDemoStop) String() string { return proto.CompactTextString(m) }
func (*CDemoStop) ProtoMessage()    {}

type CDemoUserCmd struct {
	CmdNumber        *int32 `protobuf:"varint,1,opt,name=cmd_number" json:"cmd_number,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoUserCmd) Reset()         { *m = CDemoUserCmd{} }
func (m *CDemoUserCmd) String() string { return proto.CompactTextString(m) }
func (*CDemoUserCmd) ProtoMessage()    {}

func (m *CDemoUserCmd) GetCmdNumber() int32 {
	if m != nil && m.CmdNumber != nil {
		return *m.CmdNumber
	}
	return 0
}

func (m *CDemoUserCmd) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDemoSpawnGroups struct {
	Count            *int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDemoSpawnGroups) Reset()         { *m = CDemoSpawnGroups{} }
func (m *CDemoSpawnGroups) String() string { return proto.CompactTextString(m) }
func (*CDemoSpawnGroups) ProtoMessage()    {}

func (m *CDemoSpawnGroups) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CDemoSpawnGroups) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("dota.EDemoCommands", EDemoCommands_name, EDemoCommands_value)
}
