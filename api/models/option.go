package models

type Option struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	PoolId      int    `json:"pool_id"`
}
