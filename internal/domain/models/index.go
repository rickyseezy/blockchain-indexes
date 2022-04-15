package models

import "math/big"

type Index struct {
	Name              string   `json:"name"`
	EthPriceInWei     *big.Int `json:"ethPriceInWei"`
	UsdPriceInCents   *big.Int `json:"usdPriceInCents"`
	UsdCapitalization *big.Int `json:"usdCapitalization"`
	PercentageChange  *big.Int `json:"percentageChange"`
}
