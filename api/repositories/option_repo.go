package repositories

import (
	"strconv"
	"voter/api/core/services"
	"voter/api/core/utils"
	"voter/api/models"
)

type OptionRepo interface {
	BulkCreate(options []*models.Option) error
	BulkDelete(pollId int) error
	GetByPolls(polls []*models.Poll) ([]*models.Option, error)
	GetById(id int) (*models.Option, error)
	All() ([]*models.Option, error)
}

type OptionRepoImpl struct {
	service *services.DBService
}

func ensureOptionsTableExists(s *services.DBService) {
	s.Execute(`CREATE TABLE IF NOT EXISTS options (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		poll_id INTEGER NOT NULL,
		FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
	)`)
}

func NewOptionRepo(service *services.DBService) *OptionRepoImpl {
	ensureOptionsTableExists(service)

	return &OptionRepoImpl{
		service: service,
	}
}

func (r *OptionRepoImpl) BulkCreate(options []*models.Option) error {
	query := "INSERT INTO options (description, poll_id) VALUES "
	args := []interface{}{}

	for _, option := range options {
		query += "(?, ?),"
		args = append(args, option.Description, option.PollId)
	}

	query = query[:len(query)-1]

	_, err := r.service.Execute(query, args...)

	return err
}

func (r *OptionRepoImpl) BulkDelete(pollId int) error {
	_, err := r.service.Execute("DELETE FROM options WHERE poll_id = ?", pollId)

	return err
}

func (r *OptionRepoImpl) GetByPolls(polls []*models.Poll) ([]*models.Option, error) {
	var options []*models.Option
	
	for _, poll := range polls {
		rows, err := r.service.Select("SELECT * FROM options WHERE poll_id = ?", poll.Id)
		if err != nil {
			return nil, err
		}
		
		defer rows.Close()
		for rows.Next() {
			option := &models.Option{}

			err = rows.Scan(&option.Id, &option.Description, &option.PollId)

			if err != nil {
				return nil, err
			}

			options = append(options, option)
		}
	}

	return options, nil
}

func (r *OptionRepoImpl) GetById(id int) (*models.Option, error) {
	rows, err := r.service.Select("SELECT * FROM options WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		option := &models.Option{}
	
		err = rows.Scan(&option.Id, &option.Description, &option.PollId)

		if err != nil {
			return nil, err
		}

		return option, nil
	}

	return nil, nil
}

func (r *OptionRepoImpl) All() ([]*models.Option, error) {
	rows, err := r.service.Select("SELECT * FROM options")
	if err != nil {
		return nil, err
	}

	var options []*models.Option
	
	defer rows.Close()
	for rows.Next() {
		option := &models.Option{}

		err = rows.Scan(&option.Id, &option.Description, &option.PollId)

		if err != nil {
			return nil, err
		}

		options = append(options, option)
	}

	return options, nil
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
	indexesToDelete := []int{}

	for i, option := range r.options {
		if option.PollId == pollId {
			indexesToDelete = append(indexesToDelete, i)
		}
	}

	for i := len(indexesToDelete) - 1; i >= 0; i-- {
		r.options = append(r.options[:indexesToDelete[i]], r.options[indexesToDelete[i]+1:]...)
	}

	return nil
}

func (r *FakeOptionRepo) GetByPolls(polls []*models.Poll) ([]*models.Option, error) {
	var options []*models.Option

	for _, poll := range polls {
		for _, option := range r.options {
			if option.PollId == poll.Id {
				options = append(options, option)
			}
		}
	}

	return options, nil
}

func (r *FakeOptionRepo) GetById(id int) (*models.Option, error) {
	for _, option := range r.options {
		if option.Id == id {
			return option, nil
		}
	}

	return nil, nil
}

func (r *FakeOptionRepo) All() ([]*models.Option, error) {
	return r.options, nil
}
