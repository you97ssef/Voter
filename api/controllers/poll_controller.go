package controllers

import (
	"strconv"
	"time"
	"voter/api/core/utils"
	"voter/api/dtos"
	"voter/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (ctl *Controller) CreatePoll(c *gin.Context) {
	var newPoll dtos.NewPollDTO

	if err := c.ShouldBindJSON(&newPoll); err != nil {
		BadRequest(c, "Invalid request")
		return
	}

	var userId = ctl.Server.Jwt.GetValue(c.MustGet("claims").(jwt.MapClaims), "id").(float64)

	poll := models.Poll{
		Description: newPoll.Description,
		Creator:     int(userId),
	}

	if newPoll.Private {
		code := utils.RandomString(models.CodeLength)
		poll.PrivateCode = &code
	}

	err := ctl.Repositories.PollRepo.Save(&poll)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error saving poll")
		return
	}

	var options []*models.Option

	for _, description := range newPoll.Options {
		options = append(options, &models.Option{Description: description, PollId: poll.Id})
	}

	err = ctl.Repositories.OptionRepo.BulkCreate(options)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error saving options")
		return
	}

	Ok(
		c,
		poll,
		"Poll created successfully",
	)
}

func (ctl *Controller) MyPolls(c *gin.Context) {
	var userId = ctl.Server.Jwt.GetValue(c.MustGet("claims").(jwt.MapClaims), "id").(float64)

	polls, err := ctl.Repositories.PollRepo.GetByCreator(int(userId))

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting polls")
		return
	}

	options, err := ctl.Repositories.OptionRepo.GetByPolls(polls)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting options")
		return
	}

	Ok(c, gin.H{"polls": polls, "options": options}, "")
}

func (ctl *Controller) DeletePoll(c *gin.Context) {
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
		NotFound(c, "Poll not found")
		return
	}

	var userId = ctl.Server.Jwt.GetValue(c.MustGet("claims").(jwt.MapClaims), "id").(float64)

	if int(userId) != poll.Creator {
		Forbidden(c, "You are not the creator of this poll")
		return
	}

	err = ctl.Repositories.OptionRepo.BulkDelete(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error deleting options")
		return
	}

	err = ctl.Repositories.PollRepo.Delete(poll.Id)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error deleting poll")
		return
	}

	Ok(c, nil, "Poll deleted successfully")
}

func (ctl *Controller) PublicPolls(c *gin.Context) {
	polls, err := ctl.Repositories.PollRepo.GetPublic()

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting polls")
		return
	}

	options, err := ctl.Repositories.OptionRepo.GetByPolls(polls)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error getting options")
		return
	}

	Ok(c, gin.H{"polls": polls, "options": options}, "")
}

func (ctl *Controller) FinishPoll(c *gin.Context) {
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
		NotFound(c, "Poll not found")
		return
	}

	var userId = ctl.Server.Jwt.GetValue(c.MustGet("claims").(jwt.MapClaims), "id").(float64)

	if int(userId) != poll.Creator {
		Forbidden(c, "You are not the creator of this poll")
		return
	}

	now := time.Now()
	poll.FinishedAt = &now

	err = ctl.Repositories.PollRepo.Save(poll)

	if err != nil {
		ctl.Server.Logger.Alert(err)
		Error(c, err, "Error finishing poll")
		return
	}

	Ok(c, poll, "Poll finished successfully")
}
