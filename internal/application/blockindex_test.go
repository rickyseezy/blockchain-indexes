package application

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rickyseezy/block/internal/datasources/mocks"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

var groupMockResponse = []*models.Group{
	{
		ID:      big.NewInt(int64(1)),
		Name:    "foo",
		Indexes: []*big.Int{big.NewInt(int64(2)), big.NewInt(int64(3))},
	},
	{
		ID:      big.NewInt(int64(2)),
		Name:    "bar",
		Indexes: []*big.Int{big.NewInt(int64(3)), big.NewInt(int64(4))},
	},
}

var indexMockResponse = []*models.Index{
	{
		Name:              "index - foo",
		EthPriceInWei:     big.NewInt(int64(1)),
		UsdPriceInCents:   big.NewInt(int64(2)),
		UsdCapitalization: big.NewInt(int64(3)),
		PercentageChange:  big.NewInt(int64(4)),
	},
	{
		Name:              "index - bar",
		EthPriceInWei:     big.NewInt(int64(5)),
		UsdPriceInCents:   big.NewInt(int64(6)),
		UsdCapitalization: big.NewInt(int64(7)),
		PercentageChange:  big.NewInt(int64(8)),
	},
}

var blockMockResponse = []*models.Block{
	{
		BlockNumber: big.NewInt(1),
		BlockInformation: &types.Header{
			ParentHash:  common.Hash{},
			UncleHash:   common.Hash{},
			Coinbase:    common.Address{},
			Root:        common.Hash{},
			TxHash:      common.Hash{},
			ReceiptHash: common.Hash{},
			Bloom:       types.Bloom{},
			Difficulty:  nil,
			Number:      nil,
			GasLimit:    1,
			GasUsed:     1,
			Time:        1,
			Extra:       nil,
			MixDigest:   common.Hash{},
			Nonce:       types.BlockNonce{},
			BaseFee:     nil,
		},
	},
	{
		BlockNumber: big.NewInt(2),
		BlockInformation: &types.Header{
			ParentHash:  common.Hash{},
			UncleHash:   common.Hash{},
			Coinbase:    common.Address{},
			Root:        common.Hash{},
			TxHash:      common.Hash{},
			ReceiptHash: common.Hash{},
			Bloom:       types.Bloom{},
			Difficulty:  nil,
			Number:      nil,
			GasLimit:    2,
			GasUsed:     2,
			Time:        2,
			Extra:       nil,
			MixDigest:   common.Hash{},
			Nonce:       types.BlockNonce{},
			BaseFee:     nil,
		},
	},
}

type BlockIndexTestSuite struct {
	suite.Suite
}

func (suite *BlockIndexTestSuite) TestGetGroup() {
	br := mocks.NewBlockRepository()
	gr := mocks.NewMockGroupRepository()
	ir := mocks.NewIndexRepository()
	initMock(br, gr, ir)

	blockIndexApp := NewBlockIndex(br, gr, ir)
	tests := []struct {
		title    string
		input    string
		wantErr  bool
		expected *models.Group
	}{
		{
			title:    "should get a group with a valid id having whitespace",
			input:    "1 ",
			wantErr:  false,
			expected: groupMockResponse[0],
		},
		{
			title:    "should get a group with a valid id",
			input:    "1",
			wantErr:  false,
			expected: groupMockResponse[0],
		},
	}

	for _, tt := range tests {
		group, err := blockIndexApp.GetGroup(context.Background(), tt.input)
		if tt.wantErr {
			assert.NotNil(suite.T(), err)
			assert.Nil(suite.T(), group)
			continue
		}

		assert.Equal(suite.T(), tt.expected, group)
	}
}

func (suite *BlockIndexTestSuite) TestGetIndex() {
	br := mocks.NewBlockRepository()
	gr := mocks.NewMockGroupRepository()
	ir := mocks.NewIndexRepository()
	initMock(br, gr, ir)

	blockIndexApp := NewBlockIndex(br, gr, ir)
	tests := []struct {
		title    string
		input    string
		wantErr  bool
		expected *models.Index
	}{
		{
			title:    "should get an index with a valid id having whitespace",
			input:    "1 ",
			wantErr:  false,
			expected: indexMockResponse[0],
		},
		{
			title:    "should get an index with a valid id",
			input:    "1",
			wantErr:  false,
			expected: indexMockResponse[0],
		},
	}

	for _, tt := range tests {
		index, err := blockIndexApp.GetIndex(context.Background(), tt.input)
		if tt.wantErr {
			assert.NotNil(suite.T(), err)
			assert.Nil(suite.T(), index)
			continue
		}

		assert.Equal(suite.T(), tt.expected, index)
	}
}

func (suite *BlockIndexTestSuite) TestGetBlock() {
	br := mocks.NewBlockRepository()
	gr := mocks.NewMockGroupRepository()
	ir := mocks.NewIndexRepository()
	initMock(br, gr, ir)

	blockIndexApp := NewBlockIndex(br, gr, ir)
	tests := []struct {
		title    string
		input    string
		wantErr  bool
		expected *models.Block
	}{
		{
			title:    "should get a block with a valid id having whitespace",
			input:    "1 ",
			wantErr:  false,
			expected: blockMockResponse[0],
		},
		{
			title:    "should get a block with a valid id",
			input:    "1",
			wantErr:  false,
			expected: blockMockResponse[0],
		},
		{
			title:    "should get a block by hash with a valid id",
			input:    "0x4699ed643e3451b1205a97d4413d93df2f4d7975c9c23a1827bea2a932124b10",
			wantErr:  false,
			expected: blockMockResponse[0],
		},
		{
			title:    "should get a block by hash with a valid id and whitespace",
			input:    "0x4699ed643e3451b1205a97d4413d93df2f4d7975c9c23a1827bea2a932124b10   ",
			wantErr:  false,
			expected: blockMockResponse[0],
		},
		{
			title:    "should get a block using latest",
			input:    "latest",
			wantErr:  false,
			expected: blockMockResponse[1],
		},
		{
			title:    "should get a block using latest with whitespace",
			input:    "latest  ",
			wantErr:  false,
			expected: blockMockResponse[1],
		},
	}

	for _, tt := range tests {
		block, err := blockIndexApp.GetBlock(context.Background(), tt.input)
		if tt.wantErr {
			assert.NotNil(suite.T(), err)
			assert.Nil(suite.T(), block)
			continue
		}

		assert.Equal(suite.T(), tt.expected, block)
	}
}

func initMock(br *mocks.MockBlockRepository, gr *mocks.MockGroupRepository, ir *mocks.MockIndexRepository) {
	for _, group := range groupMockResponse {
		gr.On("GetGroup", context.TODO(), group.ID).Return(group, nil)
	}

	for i, index := range indexMockResponse {
		ir.On("GetIndex", context.TODO(), big.NewInt(int64(i+1))).Return(index, nil)
	}

	for i, block := range blockMockResponse {
		mockHash := fmt.Sprintf("0x4699ed643e3451b1205a97d4413d93df2f4d7975c9c23a1827bea2a932124b1%d", i)
		br.On("GetBlockByNumber", context.TODO(), big.NewInt(int64(i+1))).Return(block, nil)
		br.On("GetBlockByHash", context.TODO(), common.HexToHash(mockHash)).Return(block, nil)
	}
	br.On("GetLatest", context.TODO()).Return(blockMockResponse[1], nil)
}

func TestBlockIndexTestSuite(t *testing.T) {
	suite.Run(t, new(BlockIndexTestSuite))
}
