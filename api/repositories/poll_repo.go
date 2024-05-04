package repositories

import (
	"strconv"
	"voter/api/core/utils"
	"voter/api/models"
)

type PollRepo interface {
	Save(poll *models.Poll) error
	GetByCreator(creator int) ([]*models.Poll, error)
	GetByPrivateCode(privateCode string) (*models.Poll, error)
	GetPublic() ([]*models.Poll, error)
	GetById(id int) (*models.Poll, error)
	Delete(id int) error
	GetByCode(code string) (*models.Poll, error)
}

type FakePollRepo struct {
	polls []*models.Poll
}

func NewFakePollRepo(count, users int) *FakePollRepo {
	var polls []*models.Poll

	for i := 0; i < count; i++ {
		var privateCode *string

		if utils.RandomBool() {
			code := utils.RandomString(models.CodeLength)
			privateCode = &code
		}

		polls = append(polls, &models.Poll{
			Id:          i + 1,
			Description: "poll " + strconv.Itoa(i),
			Creator:     utils.RandomInt(1, users),
			PrivateCode: privateCode,
		})
	}

	return &FakePollRepo{
		polls: polls,
	}
}

func (r *FakePollRepo) Save(poll *models.Poll) error {
	if poll.Id == 0 {
		var maxId int
		for _, p := range r.polls {
			if p.Id > maxId {
				maxId = p.Id
			}
		}
		poll.Id = maxId + 1
		r.polls = append(r.polls, poll)
	} else {
		for i, p := range r.polls {
			if p.Id == poll.Id {
				r.polls[i] = poll
				break
			}
		}
	}

	return nil
}

func (r *FakePollRepo) GetByCreator(creator int) ([]*models.Poll, error) {
	var polls []*models.Poll

	for _, p := range r.polls {
		if p.Creator == creator {
			polls = append(polls, p)
		}
	}

	return polls, nil
}

func (r *FakePollRepo) GetByPrivateCode(privateCode string) (*models.Poll, error) {
	for _, p := range r.polls {
		if p.PrivateCode != nil && *p.PrivateCode == privateCode {
			return p, nil
		}
	}

	return nil, nil
}

func (r *FakePollRepo) GetPublic() ([]*models.Poll, error) {
	var polls []*models.Poll

	for _, p := range r.polls {
		if p.PrivateCode == nil {
			polls = append(polls, p)
		}
	}

	return polls, nil
}

func (r *FakePollRepo) Delete(id int) error {
	for i, p := range r.polls {
		if p.Id == id {
			r.polls = append(r.polls[:i], r.polls[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *FakePollRepo) GetById(id int) (*models.Poll, error) {
	for _, p := range r.polls {
		if p.Id == id {
			return p, nil
		}
	}
	return nil, nil
}

func (r *FakePollRepo) GetByCode(code string) (*models.Poll, error) {
	for _, p := range r.polls {
		if p.PrivateCode != nil && *p.PrivateCode == code {
			return p, nil
		}
	}
	return nil, nil
}
