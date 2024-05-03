package repositories

import (
	"strconv"
	"voter/api/core/utils"
	"voter/api/models"
)

type OptionRepo interface {
	BulkCreate(options []*models.Option) error
	BulkDelete(pollId int) error
	GetByPollId(pollId int) ([]*models.Option, error)
}

type FakeOptionRepo struct {
	options []*models.Option
}

func NewFakeOptionRepo(count, polls int) *FakeOptionRepo {
	var options []*models.Option

	for i := 0; i < count; i++ {
		options = append(options, &models.Option{
			Id:          i + 1,
			Description: "option " + strconv.Itoa(i),
			PollId:      utils.RandomInt(1, polls),
		})
	}

	return &FakeOptionRepo{
		options: options,
	}
}

func (r *FakeOptionRepo) BulkCreate(options []*models.Option) error {
	var maxId int
	
	for _, option := range r.options {
		if option.Id > maxId {
			maxId = option.Id
		}
	}

	for _, option := range options {
		option.Id = maxId + 1
		r.options = append(r.options, option)
		maxId++
	}

	return nil
}

func (r *FakeOptionRepo) BulkDelete(pollId int) error {
	for i, option := range r.options {
		if option.PollId == pollId {
			r.options = append(r.options[:i], r.options[i+1:]...)
		}
	}

	return nil
}

func (r *FakeOptionRepo) GetByPollId(pollId int) ([]*models.Option, error) {
	var options []*models.Option

	for _, option := range r.options {
		if option.PollId == pollId {
			options = append(options, option)
		}
	}

	return options, nil
}
