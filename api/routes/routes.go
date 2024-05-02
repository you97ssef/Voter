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
	api := r.Server.Router.Group("", m.TestMiddleware())
	api.GET("/test", c.Test)
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
