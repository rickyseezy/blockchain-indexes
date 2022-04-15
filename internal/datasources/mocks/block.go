package mocks

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockBlockRepository struct {
	mock.Mock
}

func NewBlockRepository() *MockBlockRepository {
	return &MockBlockRepository{}
}

func (b *MockBlockRepository) GetBlockByNumber(ctx context.Context, ID *big.Int) (*models.Block, error) {
	args := b.Called(ctx, ID)
	return args.Get(0).(*models.Block), args.Error(1)
}

func (b *MockBlockRepository) GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error) {
	args := b.Called(ctx, hash)
	return args.Get(0).(*models.Block), args.Error(1)
}

func (b *MockBlockRepository) GetLatest(ctx context.Context) (*models.Block, error) {
	args := b.Called(ctx)
	return args.Get(0).(*models.Block), args.Error(1)
}
