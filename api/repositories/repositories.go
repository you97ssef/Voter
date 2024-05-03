package repositories

import "voter/api/core"

type Repositories struct {
	UserRepo UserRepo
	PollRepo PollRepo
	OptionRepo OptionRepo
}

func FakeRepositories(server *core.Server) *Repositories {
	return &Repositories{
		UserRepo: NewFakeUserRepo(100, server),
		PollRepo: NewFakePollRepo(1000, 100),
		OptionRepo: NewFakeOptionRepo(2000, 1000),
	}
}