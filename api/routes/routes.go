package routes

import (
	"time"

	"voter/api/controllers"
	"voter/api/core"
	"voter/api/middleware"

	"github.com/gin-contrib/cors"
)

type Routes struct {
	Server *core.Server
}

func (r *Routes) RegisterRoutes(c *controllers.Controller, m *middleware.Middleware) {
	api := r.Server.Router.Group("")
	connected := api.Group("", m.Connected())
	verified := api.Group("", m.Verified())

	api.GET("/test", c.Test)

	api.POST("/login", c.Login)
	api.POST("/register", c.Register)
	api.GET("/verify", c.Verify)
	api.GET("/resend-verification", c.ResendVerification)

	verified.POST("/polls", c.CreatePoll)
	connected.GET("/my-polls", c.MyPolls)
	api.GET("/public-polls", c.PublicPolls)
	connected.DELETE("/polls/:id", c.DeletePoll)
	connected.PUT("/polls/:id", c.FinishPoll)

	connected.POST("/votes", c.Vote)
	api.POST("/guest-votes", c.GuestVote)
	api.GET("/polls/:id", c.Votes)
	api.GET("/poll-by-code/:code", c.VotesByCode)
	api.GET("/validate-poll/:id", c.ValidateVotes)
	api.GET("/validate-poll-by-code/:code", c.ValidateVotesByCode)
}

func (r *Routes) RegisterCors() {
	r.Server.Router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
				AllowCredentials: true,
				AllowOriginFunc: func(origin string) bool {
					return origin == "*"
				},
				MaxAge: 12 * time.Hour,
			},
		),
	)
}
