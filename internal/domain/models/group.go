package models

import "math/big"

type Group struct {
	ID      *big.Int   `json:"id"`
	Name    string     `json:"name"`
	Indexes []*big.Int `json:"indexes"`
}
