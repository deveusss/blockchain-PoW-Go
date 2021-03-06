package blockchain

import (
	"context"

	"Deveuspow/Deveuscore/block"
	"Deveuspow/Deveuscore/transaction"
)

// Storage represents a blockchain storage
// This interface might be implemented by any embedded storage.
// You can read more about storage details here:
// https://bitcoin.stackexchange.com/questions/28168/what-are-the-keys-used-in-the-blockchain-leveldb-ie-what-are-the-keyvalue-pair
// https://bitcoin.stackexchange.com/questions/48959/why-is-bitcoin-core-using-leveldb-instead-of-redis-or-sqlite
type Storage interface {
	// ReadGenesisBlock returns the hard-coded genesis block
	ReadGenesisBlock(context.Context) (block.Block, error)
	// ReadBlockByHash returns a block by given hash
	ReadBlockByHash(context.Context, []byte) (block.Block, error)
	// ReadLastNBlocks returns last N blocks
	ReadLastNBlocks(context.Context, int) ([]block.Block, error)
	// ReadLastBlockHash returns a last known block hash
	ReadLastBlockHash(context.Context) ([]byte, error)
	// WriteBlock writes a given block to the storage
	WriteBlock(context.Context, block.Block) error
	// Close releases allocated resources
	Close(context.Context) error
	transaction.Storage
}
