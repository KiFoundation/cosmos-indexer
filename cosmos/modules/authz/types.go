package authz

import (
	"fmt"

	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"

	stdTypes "github.com/cosmos/cosmos-sdk/types"
	authzTypes "github.com/cosmos/cosmos-sdk/x/authz"
)

// Explicitly ignored messages for tx parsing purposes
const (
	MsgExec   = "/cosmos.authz.v1beta1.MsgExec"
	MsgGrant  = "/cosmos.authz.v1beta1.MsgGrant"
	MsgRevoke = "/cosmos.authz.v1beta1.MsgRevoke"
)

type WrapperMsgExec struct {
	txModule.Message
	MsgValue *authzTypes.MsgExec
}

func (sf *WrapperMsgExec) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*authzTypes.MsgExec)

	// Setting Msgs to nil as it is of types.any and will trigger error at marshalling
	sf.MsgValue.Msgs = nil

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgExec) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgExec and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        "", // Fill in with the sender address if available
		ReceiverAddress:      sf.MsgValue.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgExec
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgExec
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgExec
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgExec
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgExec) String() string {
	return fmt.Sprintf("WrapperMsgExec: Grantee=%s, Msgs=%v", sf.MsgValue.Grantee, sf.MsgValue.Msgs)
}

type WrapperMsgGrant struct {
	txModule.Message
	MsgValue *authzTypes.MsgGrant
}

func (sf *WrapperMsgGrant) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*authzTypes.MsgGrant)

	// Removing types.any field to avoid error
	sf.MsgValue.Grant.Authorization = nil

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgGrant) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgGrant and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.Granter,
		ReceiverAddress:      sf.MsgValue.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgGrant
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgGrant
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgGrant
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgGrant
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgGrant) String() string {
	return fmt.Sprintf("WrapperMsgGrant: Granter=%s, Grantee=%s, Grant=%v", sf.MsgValue.Granter, sf.MsgValue.Grantee, sf.MsgValue.Grant)
}

type WrapperMsgRevoke struct {
	txModule.Message
	MsgValue *authzTypes.MsgRevoke
}

func (sf *WrapperMsgRevoke) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*authzTypes.MsgRevoke)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgRevoke) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgRevoke and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.Granter,
		ReceiverAddress:      sf.MsgValue.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgRevoke
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgRevoke
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgRevoke
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgRevoke
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgRevoke) String() string {
	return fmt.Sprintf("WrapperMsgRevoke: Granter=%s, Grantee=%s, MsgTypeUrl=%s", sf.MsgValue.Granter, sf.MsgValue.Grantee, sf.MsgValue.MsgTypeUrl)
}
