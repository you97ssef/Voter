package repositories

import "voter/api/models"

type UserRepo interface {
	Save(user *models.User) error
	GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	Delete(id int) error
	All() ([]*models.User, error)
}

type FakeUserRepo struct {
	users []*models.User
}

func NewFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{
	}
}

func (r *FakeUserRepo) Save(user *models.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *FakeUserRepo) GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	for _, user := range r.users {
		if user.Username == usernameOrEmail || user.Email == usernameOrEmail {
			return user, nil
		}
	}
	return nil, nil
}

func (r *FakeUserRepo) Delete(id int) error {
	for i, user := range r.users {
		if user.Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *FakeUserRepo) All() ([]*models.User, error) {
	return r.users, nil
}