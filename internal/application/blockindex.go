package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/rickyseezy/block/internal/domain/repositories"
	"math/big"
	"strconv"
	"strings"
)

type BlockIndex struct {
	blockRepo repositories.Block
	groupRepo repositories.Group
	indexRepo repositories.Index
}

func NewBlockIndex(br repositories.Block, gr repositories.Group, ir repositories.Index) *BlockIndex {
	return &BlockIndex{
		blockRepo: br,
		groupRepo: gr,
		indexRepo: ir,
	}
}

func (b *BlockIndex) GetGroup(ctx context.Context, param string) (*models.Group, error) {
	parseParam := strings.TrimSpace(param)
	ok := isParamValid(parseParam)
	if !ok {
		return nil, errors.New("id param is not set")
	}
	ok = isNumber(parseParam)
	if !ok {
		return nil, errors.New("id param is invalid")
	}

	id, err := convertStringToBigInt(parseParam)
	if err != nil {
		return nil, err
	}

	group, err := b.groupRepo.GetGroup(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot get group '%d': error: %s", id, err.Error())
	}

	return group, nil
}

func (b *BlockIndex) ListGroups(ctx context.Context) ([]*models.Group, error) {
	groups, err := b.groupRepo.ListGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot list groups: error: %s", err.Error())
	}

	return groups, nil
}

func (b *BlockIndex) GetIndex(ctx context.Context, param string) (*models.Index, error) {
	parseParam := strings.TrimSpace(param)
	ok := isParamValid(parseParam)
	if !ok {
		return nil, errors.New("id param is not set")
	}
	ok = isNumber(parseParam)
	if !ok {
		return nil, errors.New("id param is invalid")
	}

	id, err := convertStringToBigInt(parseParam)
	if err != nil {
		return nil, err
	}

	index, err := b.indexRepo.GetIndex(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot get index '%d': error: %s", id, err.Error())
	}

	return index, nil
}

func (b *BlockIndex) GetBlock(ctx context.Context, searchBY string) (*models.Block, error) {
	parseParam := strings.TrimSpace(searchBY)
	if isNumber(parseParam) {
		return b.getBlockByNumber(ctx, parseParam)
	}

	if parseParam == "latest" {
		return b.getLatestBlock(ctx)
	}

	return b.getBlockByHash(ctx, parseParam)
}

func (b *BlockIndex) getBlockByNumber(ctx context.Context, nb string) (*models.Block, error) {
	id, err := convertStringToBigInt(nb)
	if err != nil {
		return nil, err
	}

	block, err := b.blockRepo.GetBlockByNumber(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot get block by number: error: %s", err.Error())
	}
	return block, nil
}

func (b *BlockIndex) getBlockByHash(ctx context.Context, hash string) (*models.Block, error) {
	block, err := b.blockRepo.GetBlockByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return nil, fmt.Errorf("cannot get block using hash value: error: %s", err.Error())
	}
	return block, nil
}

func (b *BlockIndex) getLatestBlock(ctx context.Context) (*models.Block, error) {
	block, err := b.blockRepo.GetLatest(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get latest block: error: %s", err.Error())
	}

	return block, nil
}

func isNumber(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}

	return false
}

func isParamValid(param string) bool {
	if param == "" {
		return false
	}

	return true
}

func convertStringToBigInt(v string) (*big.Int, error) {
	bigInt, ok := new(big.Int).SetString(v, 0)
	if !ok {
		return nil, fmt.Errorf("cannot value '%s' convert to big int", v)
	}

	return bigInt, nil
}
