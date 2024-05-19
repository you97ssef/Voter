package models

import "time"

type Poll struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Creator     int     `json:"user_id"`
	PrivateCode *string `json:"private_code"`
	FinishedAt  *time.Time `json:"finished_at"`
}

const CodeLength = 6
