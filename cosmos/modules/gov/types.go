package gov

import (
	"fmt"

	"github.com/DefiantLabs/cosmos-indexer/config"
	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	MsgVote           = "/cosmos.gov.v1beta1.MsgVote"
	MsgDeposit        = "/cosmos.gov.v1beta1.MsgDeposit"        // handle additional deposits to the given proposal
	MsgSubmitProposal = "/cosmos.gov.v1beta1.MsgSubmitProposal" // handle the initial deposit for the proposer
	MsgVoteWeighted   = "/cosmos.gov.v1beta1.MsgVoteWeighted"
)

type WrapperMsgSubmitProposal struct {
	txModule.Message
	CosmosMsgSubmitProposal *govTypes.MsgSubmitProposal
	CoinReceived            stdTypes.Coin
	MultiCoinsReceived      stdTypes.Coins
	DepositReceiverAddress  string
}

type WrapperMsgDeposit struct {
	txModule.Message
	CosmosMsgDeposit       *govTypes.MsgDeposit
	CoinReceived           stdTypes.Coin
	MultiCoinsReceived     stdTypes.Coins
	DepositReceiverAddress string
}

type WrapperMsgVote struct {
	txModule.Message
	MsgVote *govTypes.MsgVote
}

type WrapperMsgVoteWeighted struct {
	txModule.Message
	MsgVoteWeighted *govTypes.MsgVoteWeighted
}

func (sf *WrapperMsgSubmitProposal) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	relevantData := make([]parsingTypes.MessageRelevantInformation, len(sf.CosmosMsgSubmitProposal.InitialDeposit))

	for i, v := range sf.CosmosMsgSubmitProposal.InitialDeposit {
		var currRelevantData parsingTypes.MessageRelevantInformation
		currRelevantData.SenderAddress = sf.CosmosMsgSubmitProposal.Proposer
		currRelevantData.ReceiverAddress = sf.DepositReceiverAddress

		// Amount always seems to be an integer, float may be an extra unneeded step
		currRelevantData.AmountSent = v.Amount.BigInt()
		currRelevantData.DenominationSent = v.Denom

		// This is required since we do CSV parsing on the receiver here too
		currRelevantData.AmountReceived = v.Amount.BigInt()
		currRelevantData.DenominationReceived = v.Denom

		relevantData[i] = currRelevantData
	}

	return relevantData
}

func (sf *WrapperMsgDeposit) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	relevantData := make([]parsingTypes.MessageRelevantInformation, len(sf.CosmosMsgDeposit.Amount))

	for i, v := range sf.CosmosMsgDeposit.Amount {
		var currRelevantData parsingTypes.MessageRelevantInformation
		currRelevantData.SenderAddress = sf.CosmosMsgDeposit.Depositor
		currRelevantData.ReceiverAddress = sf.DepositReceiverAddress

		// Amount always seems to be an integer, float may be an extra unneeded step
		currRelevantData.AmountSent = v.Amount.BigInt()
		currRelevantData.DenominationSent = v.Denom

		// This is required since we do CSV parsing on the receiver here too
		currRelevantData.AmountReceived = v.Amount.BigInt()
		currRelevantData.DenominationReceived = v.Denom

		relevantData[i] = currRelevantData
	}

	return relevantData
}

func (w *WrapperMsgVote) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgVote and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgVote.Voter,
		ReceiverAddress:      "",  // Set to empty string as we don't have this data in MsgVote
		AmountSent:           nil, // Set to nil as we don't have this data in MsgVote
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgVote
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgVote
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgVote
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (w *WrapperMsgVoteWeighted) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgVoteWeighted and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        w.MsgVoteWeighted.Voter,
		ReceiverAddress:      "",  // Set to empty string as we don't have this data in MsgVoteWeighted
		AmountSent:           nil, // Set to nil as we don't have this data in MsgVoteWeighted
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgVoteWeighted
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgVoteWeighted
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgVoteWeighted
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

// Proposal with an initial deposit
func (sf *WrapperMsgSubmitProposal) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.CosmosMsgSubmitProposal = msg.(*govTypes.MsgSubmitProposal)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// If there was an initial deposit, there will be a transfer log with sender and amount
	proposerDepositedCoinsEvt := txModule.GetEventWithType(bankTypes.EventTypeTransfer, log)
	if proposerDepositedCoinsEvt == nil {
		return nil
	}

	coinsReceived := txModule.GetValueForAttribute("amount", proposerDepositedCoinsEvt)
	recipientAccount := txModule.GetValueForAttribute("recipient", proposerDepositedCoinsEvt)
	sf.DepositReceiverAddress = recipientAccount

	// This may be able to be optimized by doing one or the other
	coin, err := stdTypes.ParseCoinNormalized(coinsReceived)
	if err != nil {
		sf.MultiCoinsReceived, err = stdTypes.ParseCoinsNormalized(coinsReceived)
		if err != nil {
			config.Log.Error("Error parsing coins normalized", err)
			return err
		}
	} else {
		sf.CoinReceived = coin
	}

	// Setting types.Any to nil to avoid mashalJSON error
	sf.CosmosMsgSubmitProposal.Content = nil

	return err
}

// Additional deposit
func (sf *WrapperMsgDeposit) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.CosmosMsgDeposit = msg.(*govTypes.MsgDeposit)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// If there was an initial deposit, there will be a transfer log with sender and amount
	proposerDepositedCoinsEvt := txModule.GetEventWithType(bankTypes.EventTypeTransfer, log)
	if proposerDepositedCoinsEvt == nil {
		return nil
	}

	coinsReceived := txModule.GetValueForAttribute("amount", proposerDepositedCoinsEvt)

	// This may be able to be optimized by doing one or the other
	coin, err := stdTypes.ParseCoinNormalized(coinsReceived)
	recipientAccount := txModule.GetValueForAttribute("recipient", proposerDepositedCoinsEvt)
	sf.DepositReceiverAddress = recipientAccount

	if err != nil {
		sf.MultiCoinsReceived, err = stdTypes.ParseCoinsNormalized(coinsReceived)
		if err != nil {
			config.Log.Error("Error parsing coins normalized", err)
			return err
		}
	} else {
		sf.CoinReceived = coin
	}

	return err
}

func (w *WrapperMsgVote) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgVote = msg.(*govTypes.MsgVote)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (w *WrapperMsgVoteWeighted) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	w.Type = msgType
	w.MsgVoteWeighted = msg.(*govTypes.MsgVoteWeighted)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(w.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgDeposit) String() string {
	return fmt.Sprintf("WrapperMsgDeposit: Address %s deposited %s",
		sf.CosmosMsgDeposit.Depositor, sf.CosmosMsgDeposit.Amount)
}

func (sf *WrapperMsgSubmitProposal) String() string {
	return fmt.Sprintf("WrapperMsgDeposit: Address %s deposited %s",
		sf.CosmosMsgSubmitProposal.Proposer, sf.CosmosMsgSubmitProposal.InitialDeposit)
}

func (w *WrapperMsgVote) String() string {
	return fmt.Sprintf("WrapperMsgVote: ProposalId=%d, Voter=%s, Option=%v",
		w.MsgVote.ProposalId, w.MsgVote.Voter, w.MsgVote.Option)
}

func (w *WrapperMsgVoteWeighted) String() string {
	return fmt.Sprintf("WrapperMsgVoteWeighted: ProposalId=%d, Voter=%s, Options=%v",
		w.MsgVoteWeighted.ProposalId, w.MsgVoteWeighted.Voter, w.MsgVoteWeighted.Options)
}
