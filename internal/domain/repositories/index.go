package repositories

import (
	"context"
	"github.com/rickyseezy/block/internal/domain/models"
	"math/big"
)

type Index interface {
	GetIndex(ctx context.Context, ID *big.Int) (*models.Index, error)
}
