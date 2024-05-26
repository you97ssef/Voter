package repositories

import (
	"voter/api/core/services"
	"voter/api/core/utils"
	"voter/api/models"
)

type VoteRepo interface {
	Save(vote *models.Vote) error
	GetByPoll(pollId int) ([]*models.Vote, error)
	AlreadyVoted(pollId int, userId int) (bool, error)
	AlreadyVotedGuest(pollId int, guest string) (bool, error)
	GetLastVote(pollId int) (*models.Vote, error)
}

type VoteRepoImpl struct {
	service *services.DBService
}

func ensureVotesTableExists(s *services.DBService) {
	s.Execute(`CREATE TABLE IF NOT EXISTS votes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		option_id INTEGER NOT NULL,
		poll_id INTEGER NOT NULL,
		guest TEXT,
		timestamp INTEGER NOT NULL,
		hash TEXT NOT NULL,
		prev_hash TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (option_id) REFERENCES options(id) ON DELETE CASCADE,
		FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
	)`)
}

func NewVoteRepo(service *services.DBService) *VoteRepoImpl {
	ensureVotesTableExists(service)

	return &VoteRepoImpl{
		service: service,
	}
}

func (r *VoteRepoImpl) Save(vote *models.Vote) error {
	if vote.Id == 0 {
		row, err := r.service.Execute(`INSERT INTO votes (user_id, option_id, poll_id, guest, timestamp, hash, prev_hash) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id`,
			vote.UserId, vote.OptionId, vote.PollId, vote.Guest, vote.Timestamp, vote.Hash, vote.PrevHash)
		
		if err != nil {
			return err
		}

		id, err := row.LastInsertId()
		if err != nil {
			return err
		}

		vote.Id = int(id)
		return nil
	} else {
		_, err := r.service.Execute(`UPDATE votes SET user_id = ?, option_id = ?, poll_id = ?, guest = ?, timestamp = ?, hash = ?, prev_hash = ? WHERE id = ?`,
			vote.UserId, vote.OptionId, vote.PollId, vote.Guest, vote.Timestamp, vote.Hash, vote.PrevHash, vote.Id)
		return err
	}
}

func (r *VoteRepoImpl) GetByPoll(pollId int) ([]*models.Vote, error) {
	rows, err := r.service.Select(`SELECT id, user_id, option_id, poll_id, guest, timestamp, hash, prev_hash FROM votes WHERE poll_id = ?`, pollId)
	if err != nil {
		return nil, err
	}

	var votes []*models.Vote

	defer rows.Close()
	for rows.Next() {
		var vote models.Vote

		err := rows.Scan(&vote.Id, &vote.UserId, &vote.OptionId, &vote.PollId, &vote.Guest, &vote.Timestamp, &vote.Hash, &vote.PrevHash)
		if err != nil {
			return nil, err
		}

		votes = append(votes, &vote)
	}

	return votes, nil
}

func (r *VoteRepoImpl) AlreadyVoted(pollId int, userId int) (bool, error) {
	rows, err := r.service.Select(`SELECT id FROM votes WHERE poll_id = ? AND user_id = ?`, pollId, userId)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return rows.Next(), nil
}

func (r *VoteRepoImpl) AlreadyVotedGuest(pollId int, guest string) (bool, error) {
	rows, err := r.service.Select(`SELECT id FROM votes WHERE poll_id = ? AND guest = ?`, pollId, guest)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return rows.Next(), nil
}

func (r *VoteRepoImpl) GetLastVote(pollId int) (*models.Vote, error) {
	rows, err := r.service.Select(`SELECT id, user_id, option_id, poll_id, guest, timestamp, hash, prev_hash FROM votes WHERE poll_id = ? ORDER BY id DESC LIMIT 1`, pollId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var vote models.Vote

		err := rows.Scan(&vote.Id, &vote.UserId, &vote.OptionId, &vote.PollId, &vote.Guest, &vote.Timestamp, &vote.Hash, &vote.PrevHash)
		if err != nil {
			return nil, err
		}

		return &vote, nil
	}

	return nil, nil
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

		vote := &models.Vote{
			Id:       i + 1,
			UserId:   userId,
			OptionId: option.Id,
			PollId:   option.PollId,
			Guest:    guest,
		}

		var prevVote *models.Vote = &models.Vote{}

		for _, v := range votes {
			if v.PollId == vote.PollId && v.Id > prevVote.Id {
				prevVote = v
			}
		}

		vote.CompleteVote(prevVote)
		votes = append(votes, vote)
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

func (r *FakeVoteRepo) GetLastVote(pollId int) (*models.Vote, error) {
	var lastVote *models.Vote

	for _, vote := range r.votes {
		if vote.PollId == pollId {
			if lastVote == nil || vote.Id > lastVote.Id {
				lastVote = vote
			}
		}
	}

	return lastVote, nil
}
