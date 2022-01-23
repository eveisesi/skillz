package web

import (
	"fmt"
	"net/http"

	"github.com/eveisesi/skillz"
	"github.com/eveisesi/skillz/internal/auth"
	"github.com/eveisesi/skillz/internal/cache"
	"github.com/eveisesi/skillz/internal/user/v2"
	"github.com/eveisesi/skillz/public"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/sirupsen/logrus"
)

type Service struct {
	app         *buffalo.App
	auth        auth.API
	user        user.API
	cache       cache.PageAPI
	logger      *logrus.Logger
	renderer    *render.Engine
	activeCache bool
}

const keyAuthenticatedUser = "authenticatedUser"
const keyAuthenticatedUserID = "authenticatedUserID"

const titleSuffix = "|| Eve Is ESI || A Third Party Eve Online App"

var defaultTitle = func() (string, string) {
	return "title", fmt.Sprintf("Welcome to Skillboard.Eve %s", titleSuffix)
}

func NewService(
	env skillz.Environment,
	sessionName string,
	logger *logrus.Logger,
	cache cache.PageAPI,

	auth auth.API,
	user user.API,

	renderer *render.Engine,
) *Service {

	s := &Service{
		cache:       cache,
		auth:        auth,
		user:        user,
		renderer:    renderer,
		logger:      logger,
		activeCache: env == skillz.Production,
	}

	s.app = buffalo.New(buffalo.Options{
		Env:         env.String(),
		SessionName: sessionName,
		WorkerOff:   true,
		Addr:        "127.0.0.1:54400",
		// Logger:      logger,
	})

	s.app.Use(s.SetCurrentUser)

	s.app.Use(s.Authorize)

	var skippableHandlers = []buffalo.Handler{
		s.indexHandler,
		s.loginHandler,
		s.robotsHandler,
		s.userHandler,
	}

	// router := s.app.Muxer()
	// router.Use(s.PageCacher)

	s.app.GET("/", s.indexHandler)
	s.app.GET("/login", s.loginHandler)
	s.app.GET("/logout", s.logoutHandler)
	// s.app.GET("/robots.txt", s.robotsHandler)
	s.app.GET("/users/settings", s.userSettingsHandler)
	s.app.POST("/users/settings", s.postUserSettingsHandler)
	s.app.GET("/users/{userID}", s.userHandler)
	s.app.Middleware.Skip(s.Authorize, skippableHandlers...)
	s.app.Middleware.Skip(s.SetCurrentUser, s.robotsHandler)

	s.app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory

	return s

}

func (s *Service) Start() error {
	return s.app.Serve()
}
