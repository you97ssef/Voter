package models

type Poll struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Creator     int     `json:"user_id"`
	PrivateCode *string `json:"private_code"`
}

const CodeLength = 6
