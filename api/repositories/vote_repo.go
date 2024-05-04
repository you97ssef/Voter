package repositories

import (
	"voter/api/core/utils"
	"voter/api/models"
)

type VoteRepo interface {
	Save(vote *models.Vote) error
	GetByPoll(pollId int) ([]*models.Vote, error)
	AlreadyVoted(pollId int, userId int) (bool, error)
	AlreadyVotedGuest(pollId int, guest string) (bool, error)
}

type FakeVoteRepo struct {
	votes []*models.Vote
}

func NewFakeVoteRepo(count, users int, options []*models.Option) *FakeVoteRepo {
	var votes []*models.Vote

	for i := 0; i < count; i++ {
		var userId *int
		var guest *string
		if utils.RandomBool() {
			user := utils.RandomInt(1, users)
			userId = &user
		} else {
			user := utils.RandomString(10)
			guest = &user
		}

		option := options[utils.RandomInt(0, len(options)-1)]

		votes = append(votes, &models.Vote{
			Id:       i + 1,
			UserId:   userId,
			OptionId: option.Id,
			PollId:   option.PollId,
			Guest:    guest,
		})
	}

	return &FakeVoteRepo{
		votes: votes,
	}
}

func (r *FakeVoteRepo) Save(vote *models.Vote) error {
	if vote.Id == 0 {
		var maxId int
		for _, v := range r.votes {
			if v.Id > maxId {
				maxId = v.Id
			}
		}
		vote.Id = maxId + 1
		r.votes = append(r.votes, vote)
	} else {
		for i, v := range r.votes {
			if v.Id == vote.Id {
				r.votes[i] = vote
				break
			}
		}
	}

	return nil
}

func (r *FakeVoteRepo) GetByPoll(pollId int) ([]*models.Vote, error) {
	var votes []*models.Vote

	for _, vote := range r.votes {
		if vote.PollId == pollId {
			votes = append(votes, vote)
		}
	}

	return votes, nil
}

func (r *FakeVoteRepo) AlreadyVoted(pollId int, userId int) (bool, error) {
	for _, vote := range r.votes {
		if vote.PollId == pollId && vote.UserId != nil && *vote.UserId == userId {
			return true, nil
		}
	}

	return false, nil
}

func (r *FakeVoteRepo) AlreadyVotedGuest(pollId int, guest string) (bool, error) {
	for _, vote := range r.votes {
		if vote.PollId == pollId && vote.Guest != nil && *vote.Guest == guest {
			return true, nil
		}
	}

	return false, nil
}
