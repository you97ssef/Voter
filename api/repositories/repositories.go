package repositories

import "voter/api/core"

type Repositories struct {
	UserRepo UserRepo
	PollRepo PollRepo
	OptionRepo OptionRepo
	VoteRepo VoteRepo
}

func FakeRepositories(server *core.Server) *Repositories {
	countUsers := 100
	countPolls := 1000
	countOptions := 2000
	countVotes := 10000

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