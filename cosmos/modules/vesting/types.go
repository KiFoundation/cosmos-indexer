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
	MsgValue *vestingTypes.MsgCreateVestingAccount
}

func (sf *WrapperMsgCreateVestingAccount) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*vestingTypes.MsgCreateVestingAccount)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgCreateVestingAccount) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgCreateVestingAccount and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.FromAddress,
		ReceiverAddress:      sf.MsgValue.ToAddress,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgCreateVestingAccount
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgCreateVestingAccount
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgCreateVestingAccount
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgCreateVestingAccount
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgCreateVestingAccount) String() string {
	return fmt.Sprintf("WrapperMsgCreateVestingAccount: FromAddress=%s, ToAddress=%s, Amount=%v, EndTime=%d, Delayed=%v",
		sf.MsgValue.FromAddress, sf.MsgValue.ToAddress, sf.MsgValue.Amount,
		sf.MsgValue.EndTime, sf.MsgValue.Delayed)
}
