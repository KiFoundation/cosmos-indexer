package db

import "time"

type Block struct {
	ID     uint
	Height uint64 `gorm:"uniqueIndex"`
}

type Tx struct {
	ID              uint
	TimeStamp       time.Time
	Hash            string `gorm:"uniqueIndex"`
	Fees            string
	Code            int64
	BlockId         uint
	Block           Block
	SignerAddressId *int //int pointer allows foreign key to be null
	SignerAddress   Address
}

type Address struct {
	ID      uint
	Address string `gorm:"uniqueIndex"`
}

type Message struct {
	ID           uint
	TxId         uint
	Tx           Tx
	MessageType  string `gorm:"index"`
	MessageIndex int
}

type TaxableEvent struct {
	ID                uint
	MessageId         uint
	Message           Message
	Amount            float64
	Denomination      string
	SenderAddressId   uint `gorm:"index:idx_sender"`
	SenderAddress     Address
	ReceiverAddressId uint `gorm:"index:idx_receiver"`
	ReceiverAddress   Address
}

//Store transactions with their messages for easy database creation
type TxDBWrapper struct {
	Tx            Tx
	SignerAddress Address
	Messages      []MessageDBWrapper
}

//Store messages with their taxable events for easy database creation
type MessageDBWrapper struct {
	Message       Message
	TaxableEvents []TaxableEvent
}
