package models

type Vote struct {
	Id       int     `json:"id"`
	UserId   *int    `json:"user_id"`
	OptionId int     `json:"option_id"`
	Guest    *string `json:"guest"`
	PollId   int     `json:"poll_id"`
}
