package repositories

import (
	"strconv"
	"voter/api/core"
	"voter/api/core/services"
	"voter/api/models"
)

type UserRepo interface {
	Save(user *models.User) error
	GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	Delete(id int) error
	All() ([]*models.User, error)
}

type UserRepoImpl struct {
	service *services.DBService
}

func ensureUsersTableExists(s *services.DBService) {
	s.Execute(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		verified_at TIMESTAMP
	)`)
}

func NewUserRepo(service *services.DBService) *UserRepoImpl {
	ensureUsersTableExists(service)

	return &UserRepoImpl{
		service: service,
	}
}

func (r *UserRepoImpl) Save(user *models.User) error {
	if user.Id == 0 {
		row, err := r.service.Execute(`INSERT INTO users (name, username, email, password, verified_at) VALUES (?, ?, ?, ?, ?)`,
			user.Name, user.Username, user.Email, user.Password, user.VerifiedAt)
		
		if err != nil {
			return err
		}

		id, err := row.LastInsertId()
		if err != nil {
			return err
		}

		user.Id = int(id)
		return nil
	} else {
		_, err := r.service.Execute(`UPDATE users SET name = ?, username = ?, email = ?, password = ?, verified_at = ? WHERE id = ?`,
			user.Name, user.Username, user.Email, user.Password, user.VerifiedAt, user.Id)
		return err
	}
}

func (r *UserRepoImpl) GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	rows, err := r.service.Select(`SELECT id, name, username, email, password, verified_at FROM users WHERE username = ? OR email = ?`, usernameOrEmail, usernameOrEmail)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.VerifiedAt)
		if err != nil {
			return nil, err
		}

		return &user, nil
	}

	return nil, nil
}

func (r *UserRepoImpl) Delete(id int) error {
	_, err := r.service.Execute("DELETE FROM users WHERE id = ?", id)
	
	return err
}

func (r *UserRepoImpl) All() ([]*models.User, error) {
	rows, err := r.service.Select(`SELECT id, name, username, email, password, verified_at FROM users`)
	if err != nil {
		return nil, err
	}

	var users []*models.User

	defer rows.Close()
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.VerifiedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
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
