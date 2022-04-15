package usecases

import (
	"context"
	"github.com/rickyseezy/block/internal/domain/models"
)

type BlockIndex interface {
	GetGroup(ctx context.Context, ID string) (*models.Group, error)
	ListGroups(ctx context.Context) ([]*models.Group, error)
	GetIndex(ctx context.Context, ID string) (*models.Index, error)
	GetBlock(ctx context.Context, searchBY string) (*models.Block, error)
}
