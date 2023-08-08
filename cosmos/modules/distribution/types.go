package distribution

import (
	"fmt"

	"github.com/DefiantLabs/cosmos-indexer/config"
	parsingTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules"
	txModule "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/util"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	MsgFundCommunityPool           = "/cosmos.distribution.v1beta1.MsgFundCommunityPool"
	MsgWithdrawValidatorCommission = "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission"
	MsgWithdrawDelegatorReward     = "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
	MsgWithdrawRewards             = "withdraw-rewards"                                   // FIXME: this is used in 2 places and only 1 will work....
	MsgSetWithdrawAddress          = "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress" // An explicitly ignored msg for tx parsing purposes
)

type WrapperMsgFundCommunityPool struct {
	txModule.Message
	MsgValue  *distTypes.MsgFundCommunityPool
	Depositor string
	Funds     stdTypes.Coins
}

type WrapperMsgWithdrawValidatorCommission struct {
	txModule.Message
	MsgValue                 *distTypes.MsgWithdrawValidatorCommission
	DelegatorReceiverAddress string
	CoinsReceived            stdTypes.Coin
	MultiCoinsReceived       stdTypes.Coins
}

type WrapperMsgWithdrawDelegatorReward struct {
	txModule.Message
	MsgValue           *distTypes.MsgWithdrawDelegatorReward
	CoinReceived       stdTypes.Coin
	MultiCoinsReceived stdTypes.Coins
	RecipientAddress   string
}

// HandleMsg: Handle type checking for MsgFundCommunityPool
func (sf *WrapperMsgFundCommunityPool) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*distTypes.MsgFundCommunityPool)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// Funds sent and sender address are pulled from the parsed Cosmos Msg
	sf.Depositor = sf.MsgValue.Depositor
	sf.Funds = sf.MsgValue.Amount

	return nil
}

// HandleMsg: Handle type checking for WrapperMsgWithdrawValidatorCommission
func (sf *WrapperMsgWithdrawValidatorCommission) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*distTypes.MsgWithdrawValidatorCommission)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// The attribute in the log message that shows you the delegator withdrawal address and amount received
	delegatorReceivedCoinsEvt := txModule.GetEventWithType(distTypes.EventTypeWithdrawCommission, log)
	if delegatorReceivedCoinsEvt == nil {
		return &txModule.MessageLogFormatError{MessageType: msgType, Log: fmt.Sprintf("%+v", log)}
	}

	sf.DelegatorReceiverAddress = txModule.GetValueForAttribute(bankTypes.AttributeKeyReceiver, delegatorReceivedCoinsEvt)
	coinsReceived := txModule.GetValueForAttribute("amount", delegatorReceivedCoinsEvt)

	coin, err := stdTypes.ParseCoinNormalized(coinsReceived)
	if err != nil {
		sf.MultiCoinsReceived, err = stdTypes.ParseCoinsNormalized(coinsReceived)
		if err != nil {
			fmt.Println("Error parsing coins normalized")
			fmt.Println(err)
			return err
		}
	} else {
		sf.CoinsReceived = coin
	}

	return err
}

// CosmUnmarshal(): Unmarshal JSON for MsgWithdrawDelegatorReward
func (sf *WrapperMsgWithdrawDelegatorReward) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*distTypes.MsgWithdrawDelegatorReward)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	// The attribute in the log message that shows you the delegator withdrawal address and amount received
	delegatorReceivedCoinsEvt := txModule.GetEventWithType(bankTypes.EventTypeTransfer, log)
	if delegatorReceivedCoinsEvt == nil {
		// A withdrawal without a transfer means no amounts were actually moved.
		return nil
	}

	sf.RecipientAddress = txModule.GetValueForAttribute(bankTypes.AttributeKeyRecipient, delegatorReceivedCoinsEvt)
	coinsReceived := txModule.GetValueForAttribute("amount", delegatorReceivedCoinsEvt)

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

	return err
}

func (sf *WrapperMsgFundCommunityPool) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	relevantData := make([]parsingTypes.MessageRelevantInformation, len(sf.Funds))

	for i, v := range sf.Funds {
		relevantData[i].AmountSent = v.Amount.BigInt()
		relevantData[i].DenominationSent = v.Denom
		relevantData[i].SenderAddress = sf.Depositor
	}

	return relevantData
}

func (sf *WrapperMsgWithdrawValidatorCommission) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	if sf.CoinsReceived.IsNil() {
		relevantData := make([]parsingTypes.MessageRelevantInformation, len(sf.MultiCoinsReceived))

		for i, v := range sf.MultiCoinsReceived {
			relevantData[i] = parsingTypes.MessageRelevantInformation{
				AmountReceived:       v.Amount.BigInt(),
				DenominationReceived: v.Denom,
				SenderAddress:        "",
				ReceiverAddress:      sf.DelegatorReceiverAddress,
			}
		}

		return relevantData
	}
	relevantData := make([]parsingTypes.MessageRelevantInformation, 1)
	relevantData[0] = parsingTypes.MessageRelevantInformation{
		AmountReceived:       sf.CoinsReceived.Amount.BigInt(),
		DenominationReceived: sf.CoinsReceived.Denom,
		SenderAddress:        "",
		ReceiverAddress:      sf.DelegatorReceiverAddress,
	}
	return relevantData
}

func (sf *WrapperMsgWithdrawDelegatorReward) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	if sf.CoinReceived.IsNil() {
		relevantData := make([]parsingTypes.MessageRelevantInformation, len(sf.MultiCoinsReceived))
		for i, v := range sf.MultiCoinsReceived {
			relevantData[i] = parsingTypes.MessageRelevantInformation{
				AmountReceived:       v.Amount.BigInt(),
				DenominationReceived: v.Denom,
				SenderAddress:        "",
				ReceiverAddress:      sf.RecipientAddress,
			}
		}
		return relevantData
	}
	relevantData := make([]parsingTypes.MessageRelevantInformation, 1)
	relevantData[0] = parsingTypes.MessageRelevantInformation{
		AmountReceived:       sf.CoinReceived.Amount.BigInt(),
		DenominationReceived: sf.CoinReceived.Denom,
		SenderAddress:        "",
		ReceiverAddress:      sf.RecipientAddress,
	}
	return relevantData
}

func (sf *WrapperMsgWithdrawDelegatorReward) String() string {
	var coinsReceivedString string
	if !sf.CoinReceived.IsNil() {
		coinsReceivedString = sf.CoinReceived.String()
	} else {
		coinsReceivedString = sf.MultiCoinsReceived.String()
	}

	return fmt.Sprintf("MsgWithdrawDelegatorReward: Delegator %s received %s",
		sf.MsgValue.DelegatorAddress, coinsReceivedString)
}

func (sf *WrapperMsgWithdrawValidatorCommission) String() string {
	var coinsReceivedString string
	if !sf.CoinsReceived.IsNil() {
		coinsReceivedString = sf.CoinsReceived.String()
	} else {
		coinsReceivedString = sf.MultiCoinsReceived.String()
	}

	return fmt.Sprintf("WrapperMsgWithdrawValidatorCommission: Validator %s commission withdrawn. Delegator %s received %s",
		sf.MsgValue.ValidatorAddress, sf.DelegatorReceiverAddress, coinsReceivedString)
}

func (sf *WrapperMsgFundCommunityPool) String() string {
	coinsReceivedString := sf.MsgValue.Amount.String()
	depositorAddress := sf.MsgValue.Depositor

	return fmt.Sprintf("MsgFundCommunityPool: Depositor %s gave %s",
		depositorAddress, coinsReceivedString)
}

type WrapperMsgSetWithdrawAddress struct {
	txModule.Message
	MsgValue *distTypes.MsgSetWithdrawAddress
}

func (sf *WrapperMsgSetWithdrawAddress) HandleMsg(msgType string, msg stdTypes.Msg, log *txModule.LogMessage) error {
	sf.Type = msgType
	sf.MsgValue = msg.(*distTypes.MsgSetWithdrawAddress)

	// Confirm that the action listed in the message log matches the Message type
	validLog := txModule.IsMessageActionEquals(sf.GetType(), log)
	if !validLog {
		return util.ReturnInvalidLog(msgType, log)
	}

	return nil
}

func (sf *WrapperMsgSetWithdrawAddress) ParseRelevantData() []parsingTypes.MessageRelevantInformation {
	var relevantData []parsingTypes.MessageRelevantInformation

	// Extract data from the MsgSetWithdrawAddress and populate the relevant fields in MessageRelevantInformation struct.
	currRelevantData := parsingTypes.MessageRelevantInformation{
		SenderAddress:        sf.MsgValue.DelegatorAddress,
		ReceiverAddress:      sf.MsgValue.WithdrawAddress,
		AmountSent:           nil, // Set to nil as we don't have this data in MsgSetWithdrawAddress
		AmountReceived:       nil, // Set to nil as we don't have this data in MsgSetWithdrawAddress
		DenominationSent:     "",  // Set to empty string as we don't have this data in MsgSetWithdrawAddress
		DenominationReceived: "",  // Set to empty string as we don't have this data in MsgSetWithdrawAddress
	}

	relevantData = append(relevantData, currRelevantData)

	return relevantData
}

func (sf *WrapperMsgSetWithdrawAddress) String() string {
	return fmt.Sprintf("WrapperMsgSetWithdrawAddress: DelegatorAddress=%s, WithdrawAddress=%s",
		sf.MsgValue.DelegatorAddress, sf.MsgValue.WithdrawAddress)
}
