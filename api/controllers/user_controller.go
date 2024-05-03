package controllers

import (
	"time"
	"voter/api/dtos"
	"voter/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (ctl *Controller) Login(c *gin.Context) {
	var login dtos.LoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		BadRequest(c, "Invalid request")
		return
	}

	user, err := ctl.Repositories.UserRepo.GetByUsernameOrEmail(login.UsernameOrEmail)

	if err != nil {
		Error(c, err, "Error getting user")
		return
	}

	if user == nil {
		BadRequest(c, "Username or email are invalid")
		return
	}

	if !ctl.Server.Hasher.VerifyPassword(user.Password, login.Password) {
		BadRequest(c, "Password is invalid")
		return
	}

	jwt, err := ctl.prepareToken(user)

	if err != nil {
		Error(c, err, "Error generating token")
		return
	}

	Ok(c, jwt, "")
}

func (ctl *Controller) prepareToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id": user.Id,
		"name": user.Name,
		"username": user.Username,
		"email": user.Email,
		"verified_at": user.VerifiedAt,
	}

	return ctl.Server.Jwt.GenerateToken(claims)
}

func (ctl *Controller) Register(c *gin.Context) {
	var register dtos.RegisterDTO

	if err := c.ShouldBindJSON(&register); err != nil {
		BadRequest(c, "Invalid request")
		return
	}

	user, err := ctl.Repositories.UserRepo.GetByUsernameOrEmail(register.Username)

	if err != nil {
		Error(c, err, "Error checking username")
		return
	}

	if user != nil {
		BadRequest(c, "Username already exists")
		return
	}

	user, err = ctl.Repositories.UserRepo.GetByUsernameOrEmail(register.Email)

	if err != nil {
		Error(c, err, "Error checking email")
		return
	}

	if user != nil {
		BadRequest(c, "Email already exists")
		return
	}

	password, err := ctl.Server.Hasher.HashPassword(register.Password)

	if err != nil {
		Error(c, err, "Error hashing password")
		return
	}

	user = &models.User{
		Name:     register.Name,
		Username: register.Username,
		Email:    register.Email,
		Password: password,
	}

	err = ctl.Repositories.UserRepo.Save(user)

	if err != nil {
		Error(c, err, "Error saving user")
		return
	}

	err = ctl.sendVerification(user.Email)

	if err != nil {
		Error(c, err, "Error sending verification email")
		return
	}

	token, err := ctl.prepareToken(user)
	
	if err != nil {
		Error(c, err, "Error generating token")
		return
	}

	Created(c, token, "Account created successfully. Check your email to verify your account")
}

func (ctl *Controller) sendVerification(email string) error {
	var to []string
	to = append(to, email)

	token, err := ctl.Server.Jwt.GenerateToken(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	if err != nil {
		return err
	}

	return ctl.Server.Mailer.SendEmail(
		to,
		"Verify your email",
		"Click the link below to verify your email: http://localhost:"+ctl.Server.Config.Port+"/verify?token="+token,
	)
}

func (ctl *Controller) Verify(c *gin.Context) {
	token := c.Query("token")

	claims, err := ctl.Server.Jwt.VerifyToken(token)

	if err != nil {
		Unauthorized(c, "Invalid token")
		return
	}

	email := claims["email"].(string)

	user, err := ctl.Repositories.UserRepo.GetByUsernameOrEmail(email)

	if err != nil {
		Error(c, err, "Error getting account by email")
		return
	}

	if user == nil {
		Unauthorized(c, "Account not found")
		return
	}

	user.VerifiedAt = &time.Time{}

	err = ctl.Repositories.UserRepo.Save(user)

	if err != nil {
		Error(c, err, "Error saving account")
		return
	}

	Ok(c, nil, "Email verified successfully")
}

func (ctl *Controller) ResendVerification(c *gin.Context) {
	var email string = c.Query("email")

	if email == "" {
		BadRequest(c, "Email is required")
		return
	}

	user, err := ctl.Repositories.UserRepo.GetByUsernameOrEmail(email)

	if err != nil {
		Error(c, err, "Error getting account by email")
		return
	}

	if user == nil {
		BadRequest(c, "Email not found")
		return
	}

	err = ctl.sendVerification(user.Email)

	if err != nil {
		Error(c, err, "Error sending verification email")
		return
	}

	Ok(c, nil, "Verification email sent successfully")
}
