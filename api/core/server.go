package core

import (
	"time"

	"voter/api/core/services"
	"voter/api/core/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router     *gin.Engine
	Logger     *services.Logger
	Config     *utils.Config
	Filesystem *services.Filesystem
	Mailer     *services.Mailer
	Hasher     *services.Hasher
	Jwt        *services.Jwt
	Globals    map[string]interface{}
}

func (s *Server) Initialize(envFile string) {
	configGin()
	setupLocation()

	var environment = s.setupEnv(envFile)
	s.setupGlobals()
	s.setupLogger()
	s.setupMailer(environment)
	s.setupHasher(environment)
	s.setupJwt(environment)
	s.checkSetup()
	s.setupRouter()
}

func (s *Server) Run() {
	s.Router.Run(":" + s.Config.Port)
}

func (s *Server) SetGlobal(key string, value interface{}) {
	s.Globals[key] = value
}

func configGin() {
}

func setupLocation() {
	location, _ := time.LoadLocation(time.UTC.String())
	time.Local = location
}

func (s *Server) setupEnv(envFile string) interface{} {
	s.Filesystem = services.NewFilesystem()

	var environment interface{} = nil

	s.Filesystem.ReadJSONFromFile(envFile, &environment)

	s.Config = &utils.Config{
		Port:               environment.(map[string]interface{})["config"].(map[string]interface{})["port"].(string),
		LogFile:            environment.(map[string]interface{})["config"].(map[string]interface{})["log_file"].(string),
		DatabaseConnection: environment.(map[string]interface{})["config"].(map[string]interface{})["database_connection"].(string),
	}

	return environment
}

func (s *Server) setupGlobals() {
	s.Globals = make(map[string]interface{})
}

func (s *Server) setupLogger() {
	if s.Config.LogFile != "" {
		s.Logger = services.NewFileLogger(s.Config.LogFile)
	} else {
		s.Logger = services.NewDefaultLogger()
	}
}

func (s *Server) setupMailer(environment interface{}) {
	s.Mailer = services.NewMailer(
		environment.(map[string]interface{})["mailer"].(map[string]interface{})["username"].(string),
		environment.(map[string]interface{})["mailer"].(map[string]interface{})["password"].(string),
		environment.(map[string]interface{})["mailer"].(map[string]interface{})["host"].(string),
		environment.(map[string]interface{})["mailer"].(map[string]interface{})["port"].(string),
	)
}

func (s *Server) setupHasher(environment interface{}) {
	s.Hasher = services.NewHasher(
		int(environment.(map[string]interface{})["hasher"].(map[string]interface{})["cost"].(float64)),
	)
}

func (s *Server) setupJwt(environment interface{}) {
	s.Jwt = services.NewJwt(environment.(map[string]interface{})["jwt"].(map[string]interface{})["secret_key"].(string))
}

func (s *Server) setupRouter() {
	s.Router = gin.Default()
}

func (s *Server) checkSetup() {
	checkFile(s, s.Config.LogFile)
}

func checkFile(s *Server, file string) error {
	if !s.Filesystem.FileExists(file) {
		return s.Filesystem.CreateEmptyFile(file)
	}

	return nil
}

