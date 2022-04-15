package mocks

import (
	"context"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockGroupRepository struct {
	mock.Mock
}

func NewMockGroupRepository() *MockGroupRepository {
	return &MockGroupRepository{}
}

func (g *MockGroupRepository) ListGroups(ctx context.Context) ([]*models.Group, error) {
	args := g.Called(ctx)
	return args.Get(0).([]*models.Group), args.Error(1)
}

func (g *MockGroupRepository) GetGroup(ctx context.Context, ID *big.Int) (*models.Group, error) {
	args := g.Called(ctx, ID)
	return args.Get(0).(*models.Group), args.Error(1)
}
