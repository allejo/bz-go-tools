package networking

import "encoding/binary"

const (
	MsgAccept            = "ac"
	MsgAdminInfo         = "ai"
	MsgAlive             = "al"
	MsgAddPlayer         = "ap"
	MsgAutoPilot         = "au"
	MsgCaptureFlag       = "cf"
	MsgCustomSound       = "cs"
	MsgCacheURL          = "cu"
	MsgDropFlag          = "df"
	MsgEnter             = "en"
	MsgExit              = "ex"
	MsgFlagType          = "ft"
	MsgFlagUpdate        = "fu"
	MsgFetchResources    = "fr"
	MsgGrabFlag          = "gf"
	MsgGMUpdate          = "gm"
	MsgGetWorld          = "gw"
	MsgGameSettings      = "gs"
	MsgGameTime          = "gt"
	MsgHandicap          = "hc"
	MsgKilled            = "kl"
	MsgLagState          = "ls"
	MsgMessage           = "mg"
	MsgNearFlag          = "Nf"
	MsgNewRabbit         = "nR"
	MsgNegotiateFlags    = "nf"
	MsgPause             = "pa"
	MsgPlayerInfo        = "pb"
	MsgPlayerUpdate      = "pu"
	MsgPlayerUpdateSmall = "ps"
	MsgQueryGame         = "qg"
	MsgQueryPlayers      = "qp"
	MsgReject            = "rj"
	MsgRemovePlayer      = "rp"
	MsgReplayReset       = "rr"
	MsgShotBegin         = "sb"
	MsgScore             = "sc"
	MsgScoreOver         = "so"
	MsgShotEnd           = "se"
	MsgSuperKill         = "sk"
	MsgSetVar            = "sv"
	MsgTimeUpdate        = "to"
	MsgTeleport          = "tp"
	MsgTransferFlag      = "tf"
	MsgTeamUpdate        = "tu"
	MsgWantWHash         = "wh"
	MsgWantSettings      = "ws"
	MsgPortalAdd         = "Pa"
	MsgPortalRemove      = "Pr"
	MsgPortalUpdate      = "Pu"
)

func codeFromChars(code string) uint16 {
	bytes := []byte(code)

	return binary.BigEndian.Uint16(bytes[0:2])
}

func UnpackNetworkPacket(code uint16, data[]byte) (packet interface{}) {
	switch code {
	case codeFromChars(MsgAddPlayer):
		var p MsgAddPlayerPacket
		packet = p.Unpack(data)
		break

	case codeFromChars(MsgSetVar):
		var p MsgSetVarPacket
		packet = p.Unpack(data)
		break

	case codeFromChars(MsgTeamUpdate):
		var p MsgTeamUpdatePacket
		packet = p.Unpack(data)
		break

	case codeFromChars(MsgFlagUpdate):
		var p MsgFlagUpdatePacket
		packet = p.Unpack(data)
		break
	}

	return
}
