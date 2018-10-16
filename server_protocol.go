package main

import (
	"encoding/json"
	"fmt"
)

const (
	MsgAccept            = 0x6163 // 'ac'
	MsgAdminInfo         = 0x6169 // 'ai'
	MsgAlive             = 0x616c // 'al'
	MsgAddPlayer         = 0x6170 // 'ap'
	MsgAutoPilot         = 0x6175 // 'au'
	MsgCaptureFlag       = 0x6366 // 'cf'
	MsgCustomSound       = 0x6373 // 'cs'
	MsgCacheURL          = 0x6375 // 'cu'
	MsgDropFlag          = 0x6466 // 'df'
	MsgEnter             = 0x656e // 'en'
	MsgExit              = 0x6578 // 'ex'
	MsgFlagType          = 0x6674 // 'ft'
	MsgFlagUpdate        = 0x6675 // 'fu'
	MsgFetchResources    = 0x6672 // 'fr'
	MsgGrabFlag          = 0x6766 // 'gf'
	MsgGMUpdate          = 0x676d // 'gm'
	MsgGetWorld          = 0x6777 // 'gw'
	MsgGameSettings      = 0x6773 // 'gs'
	MsgGameTime          = 0x6774 // 'gt'
	MsgHandicap          = 0x6863 // 'hc'
	MsgKilled            = 0x6b6c // 'kl'
	MsgLagState          = 0x6c73 // 'ls'
	MsgMessage           = 0x6d67 // 'mg'
	MsgNearFlag          = 0x4e66 // 'Nf'
	MsgNewRabbit         = 0x6e52 // 'nR'
	MsgNegotiateFlags    = 0x6e66 // 'nf'
	MsgPause             = 0x7061 // 'pa'
	MsgPlayerInfo        = 0x7062 // 'pb'
	MsgPlayerUpdate      = 0x7075 // 'pu'
	MsgPlayerUpdateSmall = 0x7073 // 'ps'
	MsgQueryGame         = 0x7167 // 'qg'
	MsgQueryPlayers      = 0x7170 // 'qp'
	MsgReject            = 0x726a // 'rj'
	MsgRemovePlayer      = 0x7270 // 'rp'
	MsgReplayReset       = 0x7272 // 'rr'
	MsgShotBegin         = 0x7362 // 'sb'
	MsgScore             = 0x7363 // 'sc'
	MsgScoreOver         = 0x736f // 'so'
	MsgShotEnd           = 0x7365 // 'se'
	MsgSuperKill         = 0x736b // 'sk'
	MsgSetVar            = 0x7376 // 'sv'
	MsgTimeUpdate        = 0x746f // 'to'
	MsgTeleport          = 0x7470 // 'tp'
	MsgTransferFlag      = 0x7466 // 'tf'
	MsgTeamUpdate        = 0x7475 // 'tu'
	MsgWantWHash         = 0x7768 // 'wh'
	MsgWantSettings      = 0x7773 // 'ws'
	MsgPortalAdd         = 0x5061 // 'Pa'
	MsgPortalRemove      = 0x5072 // 'Pr'
	MsgPortalUpdate      = 0x5075 // 'Pu'
)

func packetToJson(len uint32, code uint16, data []byte) {
	switch code {
	case MsgAddPlayer:
		player := handleMsgAddPlayer(len, code, data)
		unm, err := json.Marshal(player)

		if err != nil {
			fmt.Print("Unable to read packet")
			break
		}

		fmt.Print(string(unm))
		break
	}
}
