package mocks

import (
	"context"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockIndexRepository struct {
	mock.Mock
}

func NewIndexRepository() *MockIndexRepository {
	return &MockIndexRepository{}
}

func (i *MockIndexRepository) GetIndex(ctx context.Context, ID *big.Int) (*models.Index, error) {
	args := i.Called(ctx, ID)
	return args.Get(0).(*models.Index), args.Error(1)
}
