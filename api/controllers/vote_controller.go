package controllers

import (
	"strconv"
	"voter/api/dtos"
	"voter/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (ctl *Controller) GuestVote(c *gin.Context) {
	var newVote dtos.NewGuestVoteDTO

	if err := c.ShouldBindJSON(&newVote); err != nil {
		BadRequest(c, "Invalid request")
		return
	}

	poll, err := ctl.Repositories.PollRepo.GetById(newVote.PollId)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil || poll.FinishedAt != nil || (poll.PrivateCode != nil && (newVote.PollCode == nil || *poll.PrivateCode != *newVote.PollCode)) {
		BadRequest(c, "Invalid poll")
		return
	}

	option, err := ctl.Repositories.OptionRepo.GetById(newVote.OptionId)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting option")
		return
	}

	if option == nil || option.PollId != newVote.PollId {
		BadRequest(c, "Invalid option")
		return
	}

	alreadyVoted, err := ctl.Repositories.VoteRepo.AlreadyVotedGuest(newVote.PollId, newVote.Guest)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error checking if already voted")
		return
	}

	if alreadyVoted {
		BadRequest(c, "You already voted")
		return
	}

	vote := &models.Vote{
		OptionId: newVote.OptionId,
		PollId:   newVote.PollId,
		Guest:    &newVote.Guest,
	}

	previousVote, err := ctl.Repositories.VoteRepo.GetLastVote(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting last vote")
		return
	}

	if previousVote == nil {
		previousVote = &models.Vote{}
	}

	vote.CompleteVote(previousVote)

	if err := ctl.Repositories.VoteRepo.Save(vote); err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error saving vote")
		return
	}

	Created(c, vote, "You voted successfully")
}

func (ctl *Controller) Vote(c *gin.Context) {
	var newVote dtos.NewVoteDTO

	if err := c.ShouldBindJSON(&newVote); err != nil {
		BadRequest(c, "Invalid request")
		return
	}

	poll, err := ctl.Repositories.PollRepo.GetById(newVote.PollId)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil || poll.FinishedAt != nil || (poll.PrivateCode != nil && (newVote.PollCode == nil || *poll.PrivateCode != *newVote.PollCode)) {
		BadRequest(c, "Invalid poll")
		return
	}

	option, err := ctl.Repositories.OptionRepo.GetById(newVote.OptionId)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting option")
		return
	}

	if option == nil || option.PollId != newVote.PollId {
		BadRequest(c, "Invalid option")
		return
	}

	var userId = ctl.Server.Jwt.GetValue(c.MustGet("claims").(jwt.MapClaims), "id").(float64)

	alreadyVoted, err := ctl.Repositories.VoteRepo.AlreadyVoted(newVote.PollId, int(userId))

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error checking if already voted")
		return
	}

	if alreadyVoted {
		BadRequest(c, "You already voted")
		return
	}

	vote := &models.Vote{
		OptionId: newVote.OptionId,
		PollId:   newVote.PollId,
	}

	vote.UserId = new(int)
	*vote.UserId = int(userId)

	previousVote, err := ctl.Repositories.VoteRepo.GetLastVote(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting last vote")
		return
	}

	if previousVote == nil {
		previousVote = &models.Vote{}
	}

	vote.CompleteVote(previousVote)

	if err := ctl.Repositories.VoteRepo.Save(vote); err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error saving vote")
		return
	}

	Created(c, vote, "You voted successfully")
}

func (ctl *Controller) Votes(c *gin.Context) {
	var id = c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		BadRequest(c, "Invalid poll id")
		return
	}

	poll, err := ctl.Repositories.PollRepo.GetById(idInt)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil || poll.PrivateCode != nil {
		BadRequest(c, "Invalid poll")
		return
	}

	votes, err := ctl.Repositories.VoteRepo.GetVotersByPoll(idInt)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting votes")
		return
	}

	options, err := ctl.Repositories.OptionRepo.GetByPolls([]*models.Poll{poll})

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting options")
		return
	}

	Ok(c, gin.H{"votes": votes, "poll": poll, "options": options}, "")
}

func (ctl *Controller) VotesByCode(c *gin.Context) {
	var code = c.Param("code")

	poll, err := ctl.Repositories.PollRepo.GetByPrivateCode(code)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil {
		BadRequest(c, "Invalid poll code")
		return
	}

	votes, err := ctl.Repositories.VoteRepo.GetVotersByPoll(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting votes")
		return
	}

	options, err := ctl.Repositories.OptionRepo.GetByPolls([]*models.Poll{poll})

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting options")
		return
	}

	Ok(c, gin.H{"votes": votes, "poll": poll, "options": options}, "")
}

func (ctl *Controller) ValidateVotes(c *gin.Context) {
	var id = c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		BadRequest(c, "Invalid poll id")
		return
	}

	poll, err := ctl.Repositories.PollRepo.GetById(idInt)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil {
		BadRequest(c, "Invalid poll")
		return
	}

	votes, err := ctl.Repositories.VoteRepo.GetByPoll(idInt)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting votes")
		return
	}

	var votesValidation = []dtos.VoteValidation{}

	for index, vote := range votes {
		var prevVote *models.Vote = &models.Vote{}

		if index > 0 {
			prevVote = votes[index-1]
		}

		var valid = vote.IsVoteValid(prevVote)

		votesValidation = append(votesValidation, dtos.VoteValidation{Valid: valid, Id: vote.Id})
	}

	Ok(c, votesValidation, "")
}

func (ctl *Controller) ValidateVotesByCode(c *gin.Context) {
	var code = c.Param("code")

	poll, err := ctl.Repositories.PollRepo.GetByPrivateCode(code)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting poll")
		return
	}

	if poll == nil {
		BadRequest(c, "Invalid poll code")
		return
	}

	votes, err := ctl.Repositories.VoteRepo.GetByPoll(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting votes")
		return
	}

	var votesValidation = []dtos.VoteValidation{}

	for index, vote := range votes {
		var prevVote *models.Vote = &models.Vote{}

		if index > 0 {
			prevVote = votes[index-1]
		}

		var valid = vote.IsVoteValid(prevVote)

		votesValidation = append(votesValidation, dtos.VoteValidation{Valid: valid, Id: vote.Id})
	}

	Ok(c, votesValidation, "")
}
