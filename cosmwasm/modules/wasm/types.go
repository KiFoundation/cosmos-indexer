package wasm

import (
	"fmt"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	MsgExecuteContract     = "/cosmwasm.wasm.v1.MsgExecuteContract"
	MsgInstantiateContract = "/cosmwasm.wasm.v1.MsgInstantiateContract"
)

type WrapperMsgExecuteContract struct {
	txModule.Message
	MsgValue *wasmTypes.MsgExecuteContract
}

func (sf *WrapperMsgExecuteContract) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*wasmTypes.MsgExecuteContract)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgExecuteContract) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.Sender,
		ReceiverAddress:      sf.MsgValue.Contract,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     sf.MsgValue.Funds.String(),
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgExecuteContract) String() string {
	return fmt.Sprintf("WrapperMsgExecuteContract: Sender=%s, Contract=%s, Msg=%v, Funds=%v",
		sf.MsgValue.Sender, sf.MsgValue.Contract, sf.MsgValue.Msg, sf.MsgValue.Funds.String())
}

type WrapperMsgInstantiateContract struct {
	txModule.Message
	MsgValue *wasmTypes.MsgInstantiateContract
}

func (sf *WrapperMsgInstantiateContract) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*wasmTypes.MsgInstantiateContract)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgInstantiateContract) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.Sender,
		ReceiverAddress:      sf.MsgValue.Admin,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     sf.MsgValue.Funds.String(),
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgInstantiateContract) String() string {
	return fmt.Sprintf("WrapperMsgInstantiateContract: Sender=%s, Admin=%s, CodeID=%d, Label=%s, Msg=%v, Funds=%v",
		sf.MsgValue.Sender, sf.MsgValue.Admin, sf.MsgValue.CodeID,
		sf.MsgValue.Label, sf.MsgValue.Msg, sf.MsgValue.Funds.String())
}
