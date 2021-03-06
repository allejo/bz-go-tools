package networking

import (
	"bytes"
	"encoding/binary"
)

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
	chars := []byte(code)

	return binary.BigEndian.Uint16(chars[0:2])
}

func UnpackNetworkPacket(code uint16, data []byte) (packet interface{}) {
	buf := bytes.NewBuffer(data)

	switch code {
	case codeFromChars(MsgAddPlayer):
		var p MsgAddPlayerPacket
		return p.Unpack(buf)

	case codeFromChars(MsgAdminInfo):
		var p MsgAdminInfoPacket
		return p.Unpack(buf)

	case codeFromChars(MsgAlive):
		var p MsgAlivePacket
		return p.Unpack(buf)

	case codeFromChars(MsgCaptureFlag):
		var p MsgCaptureFlagPacket
		return p.Unpack(buf)

	case codeFromChars(MsgDropFlag):
		var p MsgDropFlagPacket
		return p.Unpack(buf)

	case codeFromChars(MsgFlagUpdate):
		var p MsgFlagUpdatePacket
		return p.Unpack(buf)

	case codeFromChars(MsgGMUpdate):
		var p MsgGMUpdatePacket
		return p.Unpack(buf)

	case codeFromChars(MsgGrabFlag):
		var p MsgFlagGrabPacket
		return p.Unpack(buf)

	case codeFromChars(MsgKilled):
		var p MsgKilledPacket
		return p.Unpack(buf)

	case codeFromChars(MsgMessage):
		var p MsgMessagePacket
		return p.Unpack(buf)

	case codeFromChars(MsgNewRabbit):
		var p MsgNewRabbitPacket
		return p.Unpack(buf)

	case codeFromChars(MsgPause):
		var p MsgPausePacket
		return p.Unpack(buf)

	case codeFromChars(MsgPlayerInfo):
		var p MsgPlayerInfoPacket
		return p.Unpack(buf)

	case codeFromChars(MsgPlayerUpdate), codeFromChars(MsgPlayerUpdateSmall):
		var p MsgPlayerUpdatePacket
		return p.Unpack(buf, code)

	case codeFromChars(MsgRemovePlayer):
		var p MsgRemovePlayerPacket
		return p.Unpack(buf)

	case codeFromChars(MsgScore):
		var p MsgScorePacket
		return p.Unpack(buf)

	case codeFromChars(MsgScoreOver):
		var p MsgScoreOverPacket
		return p.Unpack(buf)

	case codeFromChars(MsgShotBegin):
		var p MsgShotBeginPacket
		return p.Unpack(buf)

	case codeFromChars(MsgShotEnd):
		var p MsgShotEndPacket
		return p.Unpack(buf)

	case codeFromChars(MsgTimeUpdate):
		var p MsgTimeUpdatePacket
		return p.Unpack(buf)

	case codeFromChars(MsgTeleport):
		var p MsgTeleportPacket
		return p.Unpack(buf)

	case codeFromChars(MsgTransferFlag):
		var p MsgTransferFlagPacket
		return p.Unpack(buf)

	case codeFromChars(MsgSetVar):
		var p MsgSetVarPacket
		return p.Unpack(buf)

	case codeFromChars(MsgTeamUpdate):
		var p MsgTeamUpdatePacket
		return p.Unpack(buf)
	}

	return
}
