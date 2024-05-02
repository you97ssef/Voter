package repositories

import "voter/api/core"

type Repositories struct {
	UserRepo UserRepo
}

func FakeRepositories(server *core.Server) *Repositories {
	return &Repositories{
		UserRepo: NewFakeUserRepo(100, server),
	}
}