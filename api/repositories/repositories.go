package repositories

import "voter/api/core"

type Repositories struct {
	UserRepo UserRepo
	PollRepo PollRepo
	OptionRepo OptionRepo
	VoteRepo VoteRepo
}

func FakeRepositories(server *core.Server) *Repositories {
	countUsers := 0
	countPolls := 0
	countOptions := 0
	countVotes := 0

	userRepo := NewFakeUserRepo(countUsers, server)
	pollRepo := NewFakePollRepo(countPolls, countUsers)
	optionRepo := NewFakeOptionRepo(countOptions, countPolls)

	options, _ := optionRepo.All()

	voteRepo := NewFakeVoteRepo(countVotes, countUsers, options)

	return &Repositories{
		UserRepo: userRepo,
		PollRepo: pollRepo,
		OptionRepo: optionRepo,
		VoteRepo: voteRepo,
	}
}

func NewRepositories(server *core.Server) *Repositories {
	userRepo := NewUserRepo(server.DatabaseService)
	pollRepo := NewPollRepo(server.DatabaseService)
	optionRepo := NewOptionRepo(server.DatabaseService)
	voteRepo := NewVoteRepo(server.DatabaseService)

	return &Repositories{
		UserRepo: userRepo,
		PollRepo: pollRepo,
		OptionRepo: optionRepo,
		VoteRepo: voteRepo,
	}
}
