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
	MsgExecuteContract *wasmTypes.MsgExecuteContract
}

func (w *WrapperMsgExecuteContract) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgExecuteContract = msg.(*wasmTypes.MsgExecuteContract)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgExecuteContract) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgExecuteContract.Sender,
		ReceiverAddress:      w.MsgExecuteContract.Contract,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     w.MsgExecuteContract.Funds.String(),
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgExecuteContract) String() string {
	return fmt.Sprintf("WrapperMsgExecuteContract: Sender=%s, Contract=%s, Msg=%v, Funds=%v",
		w.MsgExecuteContract.Sender, w.MsgExecuteContract.Contract, w.MsgExecuteContract.Msg, w.MsgExecuteContract.Funds.String())
}

type WrapperMsgInstantiateContract struct {
	txModule.Message
	MsgInstantiateContract *wasmTypes.MsgInstantiateContract
}

func (w *WrapperMsgInstantiateContract) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgInstantiateContract = msg.(*wasmTypes.MsgInstantiateContract)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgInstantiateContract) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgInstantiateContract.Sender,
		ReceiverAddress:      w.MsgInstantiateContract.Admin,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     w.MsgInstantiateContract.Funds.String(),
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgInstantiateContract) String() string {
	return fmt.Sprintf("WrapperMsgInstantiateContract: Sender=%s, Admin=%s, CodeID=%d, Label=%s, Msg=%v, Funds=%v",
		w.MsgInstantiateContract.Sender, w.MsgInstantiateContract.Admin, w.MsgInstantiateContract.CodeID,
		w.MsgInstantiateContract.Label, w.MsgInstantiateContract.Msg, w.MsgInstantiateContract.Funds.String())
}
