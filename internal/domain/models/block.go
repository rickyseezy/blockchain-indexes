package models

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Block struct {
	BlockNumber      *big.Int      `json:"blockNumber"`
	BlockInformation *types.Header `json:"blockInformation"`
}
