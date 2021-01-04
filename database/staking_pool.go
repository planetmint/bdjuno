package database

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// SaveStakingPool allows to save for the given height the given stakingtypes pool
func (db *BigDipperDb) SaveStakingPool(pool stakingtypes.Pool, height int64) error {
	stmt := `
INSERT INTO staking_pool_history (height, bonded_tokens, not_bonded_tokens) 
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, height, pool.BondedTokens.Int64(), pool.NotBondedTokens.Int64())
	return err
}