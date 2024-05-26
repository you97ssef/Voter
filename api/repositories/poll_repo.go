package repositories

import (
	"strconv"
	"time"
	"voter/api/core/services"
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

type PollRepoImpl struct {
	service *services.DBService
}

func ensurePollsTableExists(s *services.DBService) {
	s.Execute(`CREATE TABLE IF NOT EXISTS polls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		creator INTEGER NOT NULL,
		private_code TEXT,
		finished_at TIMESTAMP
	)`)
}

func NewPollRepo(service *services.DBService) *PollRepoImpl {
	ensurePollsTableExists(service)

	return &PollRepoImpl{
		service: service,
	}
}

func (r *PollRepoImpl) Save(poll *models.Poll) error {
	if poll.Id == 0 {
		row, err := r.service.Execute(`INSERT INTO polls (description, creator, private_code, finished_at) VALUES (?, ?, ?, ?) RETURNING id`,
			poll.Description, poll.Creator, poll.PrivateCode, poll.FinishedAt)
		
		if err != nil {
			return err
		}

		id, err := row.LastInsertId()
		if err != nil {
			return err
		}

		poll.Id = int(id)
		return nil
	} else {
		_, err := r.service.Execute(`UPDATE polls SET description = ?, creator = ?, private_code = ?, finished_at = ? WHERE id = ?`,
			poll.Description, poll.Creator, poll.PrivateCode, poll.FinishedAt, poll.Id)
		return err
	}
}

func (r *PollRepoImpl) GetByCreator(creator int) ([]*models.Poll, error) {
	rows, err := r.service.Select(`SELECT id, description, creator, private_code, finished_at FROM polls WHERE creator = ?`, creator)
	if err != nil {
		return nil, err
	}

	var polls []*models.Poll

	defer rows.Close()
	for rows.Next() {
		var poll models.Poll

		err := rows.Scan(&poll.Id, &poll.Description, &poll.Creator, &poll.PrivateCode, &poll.FinishedAt)
		if err != nil {
			return nil, err
		}

		polls = append(polls, &poll)
	}

	return polls, nil
}

func (r *PollRepoImpl) GetByPrivateCode(privateCode string) (*models.Poll, error) {
	rows, err := r.service.Select(`SELECT id, description, creator, private_code, finished_at FROM polls WHERE private_code = ?`, privateCode)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var poll models.Poll

		err := rows.Scan(&poll.Id, &poll.Description, &poll.Creator, &poll.PrivateCode, &poll.FinishedAt)
		if err != nil {
			return nil, err
		}

		return &poll, nil
	}

	return nil, nil
}

func (r *PollRepoImpl) GetPublic() ([]*models.Poll, error) {
	rows, err := r.service.Select(`SELECT id, description, creator, private_code, finished_at FROM polls WHERE private_code IS NULL`)
	if err != nil {
		return nil, err
	}

	var polls []*models.Poll

	defer rows.Close()
	for rows.Next() {
		var poll models.Poll

		err := rows.Scan(&poll.Id, &poll.Description, &poll.Creator, &poll.PrivateCode, &poll.FinishedAt)
		if err != nil {
			return nil, err
		}

		polls = append(polls, &poll)
	}

	return polls, nil
}

func (r *PollRepoImpl) GetById(id int) (*models.Poll, error) {
	rows, err := r.service.Select(`SELECT id, description, creator, private_code, finished_at FROM polls WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var poll models.Poll

		err := rows.Scan(&poll.Id, &poll.Description, &poll.Creator, &poll.PrivateCode, &poll.FinishedAt)
		if err != nil {
			return nil, err
		}

		return &poll, nil
	}

	return nil, nil
}

func (r *PollRepoImpl) Delete(id int) error {
	_, err := r.service.Execute(`DELETE FROM polls WHERE id = ?`, id)
	return err
}

func (r *PollRepoImpl) GetByCode(code string) (*models.Poll, error) {
	rows, err := r.service.Select(`SELECT id, description, creator, private_code, finished_at FROM polls WHERE private_code = ?`, code)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var poll models.Poll

		err := rows.Scan(&poll.Id, &poll.Description, &poll.Creator, &poll.PrivateCode, &poll.FinishedAt)
		if err != nil {
			return nil, err
		}

		return &poll, nil
	}

	return nil, nil
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

		var finishedAt *time.Time
		if utils.RandomBool() {
			finished := utils.RandomDate(time.Now().AddDate(-5, 0, 0), time.Now())
			finishedAt = &finished
		}

		polls = append(polls, &models.Poll{
			Id:          i + 1,
			Description: "poll " + strconv.Itoa(i),
			Creator:     utils.RandomInt(1, users),
			PrivateCode: privateCode,
			FinishedAt:  finishedAt,
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
