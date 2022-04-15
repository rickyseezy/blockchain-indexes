package remote

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rickyseezy/block/internal/domain/models"
	"github.com/rickyseezy/block/pkg/abi"
	"log"
	"math/big"
	"sync"
)

type GroupRepository struct {
	contract *abi.Contract
}

func NewGroupRepository(c *abi.Contract) *GroupRepository {
	return &GroupRepository{
		contract: c,
	}
}

func (g *GroupRepository) ListGroups(ctx context.Context) ([]*models.Group, error) {
	groups := make([]*models.Group, 0)
	wg := sync.WaitGroup{}
	opts := bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	ids, err := g.contract.GetGroupIds(&opts)
	ch := make(chan *models.Group, len(ids))
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		wg.Add(1)
		go func(blockID *big.Int) {
			defer wg.Done()
			group, err := g.GetGroup(ctx, blockID)
			if err != nil {
				log.Println(fmt.Sprintf("could not get block: %v, %s", blockID, err.Error()))
				return
			}
			ch <- group
		}(id)
	}
	wg.Wait()
	close(ch)

	for group := range ch {
		groups = append(groups, group)
	}

	return groups, nil
}

func (g *GroupRepository) GetGroup(ctx context.Context, ID *big.Int) (*models.Group, error) {
	opts := bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	data, err := g.contract.GetGroup(&opts, ID)
	if err != nil {
		log.Println(fmt.Sprintf("could not get group: %v, %s", ID, err.Error()))
		return nil, err
	}
	group := models.Group{
		ID:      ID,
		Name:    data.Name,
		Indexes: data.Indexes,
	}

	return &group, nil
}
