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
	MsgExec *authzTypes.MsgExec
}

func (w *WrapperMsgExec) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgExec = msg.(*authzTypes.MsgExec)

	// Setting Msgs to nil as it is of types.any and will trigger error at marshalling
	w.MsgExec.Msgs = nil

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgExec) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgExec and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        "", // Fill in with the sender address if available
		ReceiverAddress:      w.MsgExec.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgExec
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgExec
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgExec
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgExec
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgExec) String() string {
	return fmt.Sprintf("WrapperMsgExec: Grantee=%s, Msgs=%v", w.MsgExec.Grantee, w.MsgExec.Msgs)
}

type WrapperMsgGrant struct {
	txModule.Message
	MsgGrant *authzTypes.MsgGrant
}

func (w *WrapperMsgGrant) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgGrant = msg.(*authzTypes.MsgGrant)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgGrant) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgGrant and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgGrant.Granter,
		ReceiverAddress:      w.MsgGrant.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgGrant
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgGrant
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgGrant
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgGrant
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgGrant) String() string {
	return fmt.Sprintf("WrapperMsgGrant: Granter=%s, Grantee=%s, Grant=%v", w.MsgGrant.Granter, w.MsgGrant.Grantee, w.MsgGrant.Grant)
}

type WrapperMsgRevoke struct {
	txModule.Message
	MsgRevoke *authzTypes.MsgRevoke
}

func (w *WrapperMsgRevoke) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgRevoke = msg.(*authzTypes.MsgRevoke)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgRevoke) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgRevoke and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgRevoke.Granter,
		ReceiverAddress:      w.MsgRevoke.Grantee,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgRevoke
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgRevoke
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgRevoke
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgRevoke
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgRevoke) String() string {
	return fmt.Sprintf("WrapperMsgRevoke: Granter=%s, Grantee=%s, MsgTypeUrl=%s", w.MsgRevoke.Granter, w.MsgRevoke.Grantee, w.MsgRevoke.MsgTypeUrl)
}
