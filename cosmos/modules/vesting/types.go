package vesting

import (
	"fmt"

	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"

	stdTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

// Explicitly ignored messages for tx parsing purposes
const (
	MsgCreateVestingAccount = "/cosmos.vesting.v1beta1.MsgCreateVestingAccount"
)

type WrapperMsgCreateVestingAccount struct {
	txModule.Message
	MsgCreateVestingAccount *vestingTypes.MsgCreateVestingAccount
}

func (w *WrapperMsgCreateVestingAccount) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgCreateVestingAccount = msg.(*vestingTypes.MsgCreateVestingAccount)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgCreateVestingAccount) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgCreateVestingAccount and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgCreateVestingAccount.FromAddress,
		ReceiverAddress:      w.MsgCreateVestingAccount.ToAddress,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgCreateVestingAccount
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgCreateVestingAccount
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgCreateVestingAccount
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgCreateVestingAccount
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgCreateVestingAccount) String() string {
	return fmt.Sprintf("WrapperMsgCreateVestingAccount: FromAddress=%s, ToAddress=%s, Amount=%v, EndTime=%d, Delayed=%v",
		w.MsgCreateVestingAccount.FromAddress, w.MsgCreateVestingAccount.ToAddress, w.MsgCreateVestingAccount.Amount,
		w.MsgCreateVestingAccount.EndTime, w.MsgCreateVestingAccount.Delayed)
}
