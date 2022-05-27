package network

import "Deveuspow/Deveuscore/network/message"

const (
	messageAddTx    = "add_tx"
	messageAddBlock = "add_block"
)

// NewAddTXMessage creates a new TX message
func NewAddTXMessage(txData []byte) *message.Msg {
	return &message.Msg{
		Type: messageAddTx,
		Data: txData,
	}
}

// NewAddBlock creates a new Block message
func NewAddBlock(blockData []byte) *message.Msg {
	return &message.Msg{
		Type: messageAddBlock,
		Data: blockData,
	}
}
