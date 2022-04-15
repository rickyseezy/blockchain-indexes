package repositories

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rickyseezy/block/internal/domain/models"
	"math/big"
)

type Block interface {
	GetBlockByNumber(ctx context.Context, ID *big.Int) (*models.Block, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error)
	GetLatest(ctx context.Context) (*models.Block, error)
}
