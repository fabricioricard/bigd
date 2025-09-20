package votes

import (
	"github.com/pkt-cash/pktd/txscript/opcode"
	"github.com/pkt-cash/pktd/txscript/parsescript"
)

const (
	VOTE      byte = 0x00
	CANDIDATE byte = 0x01
)

type NsVote struct {
	VoterPkScript           []byte
	VoterIsWillingCandidate bool
	VoteCastInBlock         uint32
	VoteForPkScript         []byte
}

const EpochBlocks = 60 * 24 * 7

const VoteExpirationEpochs = 52

const VoteExpirationBlocks = VoteExpirationEpochs * EpochBlocks

func GetVote(outputScript []byte) *NsVote {
	scr, err := parsescript.ParseScript(outputScript)
	if err != nil {
		return nil
	}
	if len(scr) < 1 || scr[0].Opcode.Value != opcode.OP_RETURN {
		// Normal script, does not begin with OP_RETURN
		return nil
	}
	if len(scr) < 2 || scr[1].Opcode.Value > opcode.OP_16 {
		// It's an op-return script which contains something other than a push
		return nil
	}
	if len(scr) > 2 {
		// it's an op-return script but it contains additional data after the push
		return nil
	}
	data := scr[1].Data
	if len(data) < 1 || (data[0] != VOTE && data[0] != CANDIDATE) {
		// Not a vote operation
		return nil
	}
	return &NsVote{
		VoterIsWillingCandidate: data[0] == CANDIDATE,
		VoteForPkScript:         data[1:],
	}
}
