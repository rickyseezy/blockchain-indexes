package remote

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/rickyseezy/block/pkg/abi"
	"math/big"
)

type IndexRepository struct {
	contract *abi.Contract
}

func NewIndexRepository(c *abi.Contract) *IndexRepository {
	return &IndexRepository{
		contract: c,
	}
}

func (i *IndexRepository) GetIndex(ctx context.Context, ID *big.Int) (*models.Index, error) {
	opts := bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	data, err := i.contract.GetIndex(&opts, ID)
	if err != nil {
		return nil, err
	}

	return mapDataToIndex(data), nil
}

func mapDataToIndex(data struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}) *models.Index {
	return &models.Index{
		Name:              data.Name,
		EthPriceInWei:     data.EthPriceInWei,
		UsdPriceInCents:   data.UsdPriceInCents,
		UsdCapitalization: data.UsdCapitalization,
		PercentageChange:  data.PercentageChange,
	}
}
