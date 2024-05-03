package dtos

type NewPollDTO struct {
	Description string   `json:"description" binding:"required"`
	Private     bool     `json:"private"`
	Options     []string `json:"options" binding:"required,min=2"`
}
