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

func (m *Middleware) Verified() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := m.Server.Jwt.VerifyTokenFromGinContext(c)

		if err != nil {
			controllers.Unauthorized(c, "Token not provided or invalid")
			return
		}

		verifiedAt := m.Server.Jwt.GetValue(claims, "verified_at")

		if verifiedAt == nil {
			controllers.Unauthorized(c, "Your account is not verified, please verify it first")
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
