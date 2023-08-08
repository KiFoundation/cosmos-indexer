package ibc

import (
	"fmt"

	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	trantypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clitypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
)

const (
	MsgRecvPacket      = "/ibc.core.channel.v1.MsgRecvPacket"
	MsgAcknowledgement = "/ibc.core.channel.v1.MsgAcknowledgement"

	// Explicitly ignored messages for tx parsing purposes
	MsgTransfer           = "/ibc.applications.transfer.v1.MsgTransfer"
	MsgChannelOpenTry     = "/ibc.core.channel.v1.MsgChannelOpenTry"
	MsgChannelOpenConfirm = "/ibc.core.channel.v1.MsgChannelOpenConfirm"
	MsgChannelOpenInit    = "/ibc.core.channel.v1.MsgChannelOpenInit"
	MsgChannelOpenAck     = "/ibc.core.channel.v1.MsgChannelOpenAck"

	MsgTimeout        = "/ibc.core.channel.v1.MsgTimeout"
	MsgTimeoutOnClose = "/ibc.core.channel.v1.MsgTimeoutOnClose"

	MsgConnectionOpenTry     = "/ibc.core.connection.v1.MsgConnectionOpenTry"
	MsgConnectionOpenConfirm = "/ibc.core.connection.v1.MsgConnectionOpenConfirm"
	MsgConnectionOpenInit    = "/ibc.core.connection.v1.MsgConnectionOpenInit"
	MsgConnectionOpenAck     = "/ibc.core.connection.v1.MsgConnectionOpenAck"

	MsgCreateClient = "/ibc.core.client.v1.MsgCreateClient"
	MsgUpdateClient = "/ibc.core.client.v1.MsgUpdateClient"

	// Consts used for classifying Ack messages
	// We may need to keep extending these consts for other types
	AckFungibleTokenTransfer    = 0
	AckNotFungibleTokenTransfer = 1

	// Same as above, we may want to to extend these to track other results
	AckSuccess = 0
	AckFailure = 1
)

type WrapperMsgRecvPacket struct {
	txModule.Message
	MsgValue        *chantypes.MsgRecvPacket
	Sequence        uint64
	SenderAddress   string
	ReceiverAddress string
	Amount          stdTypes.Int
	Denom           string
}

func (sf *WrapperMsgRecvPacket) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*chantypes.MsgRecvPacket)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// Unmarshal the json encoded packet data so we can access sender, receiver and denom info
	var data types.FungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(sf.MsgValue.Packet.GetData(), &data); err != nil {
		// If there was a failure then this recv was not for a token transfer packet,
		// currently we only consider successful token transfers taxable events.
		return nil
	}

	sf.SenderAddress = data.Sender
	sf.ReceiverAddress = data.Receiver
	sf.Sequence = sf.MsgValue.Packet.Sequence

	amount, ok := stdTypes.NewIntFromString(data.Amount)
	if !ok {
		return fmt.Errorf("failed to convert denom amount to sdk.Int, got(%s)", data.Amount)
	}

	sf.Amount = amount
	sf.Denom = data.Denom

	return nil
}

func (sf *WrapperMsgRecvPacket) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	// This prevents the item from being indexed
	if sf.Amount.IsNil() {
		return nil
	}

	// MsgRecvPacket indicates a user has received assets on this chain so amount sent will always be 0
	amountSent := stdTypes.NewInt(0)

	return []parsingTypes.MessageRelevantInformation{{
		SenderAddress:        sf.SenderAddress,
		ReceiverAddress:      sf.ReceiverAddress,
		AmountSent:           amountSent.BigInt(),
		AmountReceived:       sf.Amount.BigInt(),
		DenominationSent:     "",
		DenominationReceived: sf.Denom,
	}}
}

func (sf *WrapperMsgRecvPacket) String() string {
	if sf.Amount.IsNil() {
		return "MsgRecvPacket: IBC transfer was not a FungibleTokenTransfer"
	}
	return fmt.Sprintf("MsgRecvPacket: IBC transfer of %s%s from %s to %s", sf.Amount, sf.Denom, sf.SenderAddress, sf.ReceiverAddress)
}

type WrapperMsgAcknowledgement struct {
	txModule.Message
	MsgValue        *chantypes.MsgAcknowledgement
	Sequence        uint64
	SenderAddress   string
	ReceiverAddress string
	Amount          stdTypes.Int
	Denom           string
	AckType         int
	AckResult       int
}

func (sf *WrapperMsgAcknowledgement) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*chantypes.MsgAcknowledgement)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// Unmarshal the json encoded packet data so we can access sender, receiver and denom info
	var data types.FungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(sf.MsgValue.Packet.GetData(), &data); err != nil {
		// If there was a failure then this ack was not for a token transfer packet,
		// currently we only consider successful token transfers taxable events.
		sf.AckType = AckNotFungibleTokenTransfer
		return nil
	}

	sf.AckType = AckFungibleTokenTransfer

	sf.SenderAddress = data.Sender
	sf.ReceiverAddress = data.Receiver
	sf.Sequence = sf.MsgValue.Packet.Sequence

	amount, ok := stdTypes.NewIntFromString(data.Amount)
	if !ok {
		return fmt.Errorf("failed to convert denom amount to sdk.Int, got(%s)", data.Amount)
	}

	// Acknowledgements can contain an error & we only want to index successful acks,
	// so we need to check the ack bytes to determine if it was a result or an error.
	var ack chantypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(sf.MsgValue.Acknowledgement, &ack); err != nil {
		return fmt.Errorf("cannot unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	switch ack.Response.(type) {
	case *chantypes.Acknowledgement_Error:
		// We index nothing on Acknowledgement errors
		sf.AckResult = AckFailure
		return nil
	default:
		// the acknowledgement succeeded on the receiving chain
		sf.AckResult = AckSuccess
		sf.Amount = amount
		sf.Denom = data.Denom
		return nil
	}
}

func (sf *WrapperMsgAcknowledgement) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	// This prevents the item from being indexed
	if sf.Amount.IsNil() || sf.AckType == AckNotFungibleTokenTransfer || sf.AckResult == AckFailure {
		return nil
	}

	// MsgAcknowledgement indicates a user has successfully sent a packet
	// so the received amount will always be zero
	amountReceived := stdTypes.NewInt(0)

	return []parsingTypes.MessageRelevantInformation{{
		SenderAddress:        sf.SenderAddress,
		ReceiverAddress:      sf.ReceiverAddress,
		AmountSent:           sf.Amount.BigInt(),
		AmountReceived:       amountReceived.BigInt(),
		DenominationSent:     sf.Denom,
		DenominationReceived: "",
	}}
}

func (w *WrapperMsgAcknowledgement) String() string {
	if w.AckType == AckNotFungibleTokenTransfer {
		return "MsgAcknowledgement: IBC transfer was not a FungibleTokenTransfer"
	}

	if w.AckType == AckFungibleTokenTransfer && w.AckResult == AckFailure {
		return "MsgAcknowledgement: IBC transfer was not successful"
	}

	if w.Amount.IsNil() {
		return "MsgAcknowledgement: IBC transfer was not a FungibleTokenTransfer"
	}

	return fmt.Sprintf("MsgAcknowledgement: IBC transfer of %s%s from %s to %s\n", w.Amount, w.Denom, w.SenderAddress, w.ReceiverAddress)
}

type WrapperMsgUpdateClient struct {
	txModule.Message
	MsgValue *clitypes.MsgUpdateClient
}

func (sf *WrapperMsgUpdateClient) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*clitypes.MsgUpdateClient)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// Set field Header of type types.any to nil
	// elseway, it triggers an error while marshalling to json
	sf.MsgValue.Header = nil

	return nil
}

func (sf *WrapperMsgUpdateClient) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.ClientId,
		ReceiverAddress:      sf.MsgValue.Signer,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     "",
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgUpdateClient) String() string {
	return fmt.Sprintf("WrapperMsgUpdateClient: ClientID=%s, SignerAddress=%s", sf.MsgValue.ClientId, sf.MsgValue.Signer)
}

type WrapperMsgTransfer struct {
	txModule.Message
	MsgValue *trantypes.MsgTransfer
}

func (sf *WrapperMsgTransfer) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*trantypes.MsgTransfer)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgTransfer) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.Sender,
		ReceiverAddress:      sf.MsgValue.Receiver,
		AmountSent:           nil,
		AmountReceived:       nil,
		DenominationSent:     sf.MsgValue.Token.Denom,
		DenominationReceived: "",
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgTransfer) String() string {
	return fmt.Sprintf("WrapperMsgTransfer: SourcePort=%s, SourceChannel=%s, Token=%v, Sender=%s, Receiver=%s, TimeoutHeight=%v, TimeoutTimestamp=%d, Memo=%s",
		sf.MsgValue.SourcePort, sf.MsgValue.SourceChannel, sf.MsgValue.Token, sf.MsgValue.Sender, sf.MsgValue.Receiver, sf.MsgValue.TimeoutHeight, sf.MsgValue.TimeoutTimestamp, sf.MsgValue.Memo)
}
