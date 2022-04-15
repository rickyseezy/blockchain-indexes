package repositories

import (
	"context"
	"github.com/rickyseezy/block/internal/domain/models"
	"math/big"
)

type Group interface {
	ListGroups(ctx context.Context) ([]*models.Group, error)
	GetGroup(ctx context.Context, ID *big.Int) (*models.Group, error)
}
