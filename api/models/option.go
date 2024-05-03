package models

type Option struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	PollId      int    `json:"poll_id"`
}
