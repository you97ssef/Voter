package middleware

import (
	"voter/api/controllers"
	"voter/api/core"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Server *core.Server
}

func (m *Middleware) Connected() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := m.Server.Jwt.VerifyTokenFromGinContext(c)

		if err != nil {
			controllers.Unauthorized(c, "Token not provided or invalid")
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
