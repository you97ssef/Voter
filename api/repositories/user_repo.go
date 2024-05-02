package repositories

import (
	"strconv"
	"voter/api/core"
	"voter/api/models"
)

type UserRepo interface {
	Save(user *models.User) error
	GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	Delete(id int) error
	All() ([]*models.User, error)
}

type FakeUserRepo struct {
	users []*models.User
}

func NewFakeUserRepo(count int, server *core.Server) *FakeUserRepo {
	var users []*models.User

	for i := 0; i < count; i++ {
		password, err := server.Hasher.HashPassword("password" + strconv.Itoa(i))

		if err != nil {
			panic(err)
		}

		users = append(users, &models.User{
			Id:         i + 1,
			Name:       "user " + strconv.Itoa(i),
			Username:   "user" + strconv.Itoa(i),
			Email:      "user" + strconv.Itoa(i) + "@example.com",
			VerifiedAt: nil,
			Password:   password,
		})
	}

	return &FakeUserRepo{
		users: users,
	}
}

func (r *FakeUserRepo) Save(user *models.User) error {
	if user.Id == 0 {
		var maxId int
		for _, u := range r.users {
			if u.Id > maxId {
				maxId = u.Id
			}
		}
		user.Id = maxId + 1
		r.users = append(r.users, user)
	} else {
		for i, u := range r.users {
			if u.Id == user.Id {
				r.users[i] = user
				break
			}
		}
	}

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
