package yanzal

import (
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

// Yanzal is the proof-of-stake consensus engine proposed
type Yanzal struct {
	config *params.YanzalConfig
}

// New creates a Yanzal proof-of-stake consensus engine
func New(config *params.YanzalConfig, db ethdb.Database) *Yanzal {

}
