package remote

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rickyseezy/block/internal/domain/models"
	"math/big"
)

type BlockRepository struct {
	client *ethclient.Client
}

func NewBlockRepository(cl *ethclient.Client) *BlockRepository {
	return &BlockRepository{
		client: cl,
	}
}

func (b *BlockRepository) GetBlockByNumber(ctx context.Context, ID *big.Int) (*models.Block, error) {
	block, err := b.client.BlockByNumber(ctx, ID)
	if err != nil {
		return nil, err
	}

	return mapDataToBlock(block.Header(), block.Number()), nil
}

func (b *BlockRepository) GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error) {
	block, err := b.client.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	return mapDataToBlock(block.Header(), block.Number()), nil
}
func (b *BlockRepository) GetLatest(ctx context.Context) (*models.Block, error) {
	id, err := b.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	block, err := b.GetBlockByNumber(ctx, big.NewInt(int64(id)))
	if err != nil {
		return nil, err
	}

	return block, nil
}

func mapDataToBlock(blockHeader *types.Header, ID *big.Int) *models.Block {
	return &models.Block{
		BlockNumber:      ID,
		BlockInformation: blockHeader,
	}
}
