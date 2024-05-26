package main

import (
	"voter/api/controllers"
	"voter/api/core"
	"voter/api/middleware"
	"voter/api/repositories"
	"voter/api/routes"
)

func main() {
	server := &core.Server{}

	server.Initialize("./environment.json")

	repositories := repositories.NewRepositories(server)

	controllers := &controllers.Controller{
		Server: server,
		Repositories: repositories,
	}

	routes := &routes.Routes{
		Server: server,
	}

	middleware := &middleware.Middleware{
		Server: server,
	}

	routes.RegisterCors()
	routes.RegisterRoutes(controllers, middleware)

	server.SetGlobal("frontend_url", "http://localhost:5173")

	server.Run()
}
