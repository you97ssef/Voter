package dtos

type NewVoteDTO struct {
	OptionId int     `json:"option_id" binding:"required"`
	PollId   int     `json:"poll_id" binding:"required"`
	PollCode *string `json:"poll_code"`
}

type NewGuestVoteDTO struct {
	OptionId int     `json:"option_id" binding:"required"`
	PollId   int     `json:"poll_id" binding:"required"`
	Guest    string  `json:"guest" binding:"required"`
	PollCode *string `json:"poll_code"`
}
