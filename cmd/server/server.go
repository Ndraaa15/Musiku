package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Ndraaa15/musiku/internal/api/controller"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/domain/service"
	"github.com/Ndraaa15/musiku/internal/infrastructure/postgresql"
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type server struct {
	router *gin.Engine
	server *http.Server
	ctrl   *controller.Controller
}

func New() (*server, error) {
	s := &server{
		router: gin.Default(),
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	db, err := postgresql.NewPostgreSqlClient()

	if err != nil {
		log.Printf("[musiku-server] failed to initialize musiku database : %v\n", err)
		return nil, err
	}

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)

	s.router = gin.Default()
	s.ctrl = controller.NewController(us)

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()
	s.router.Run()

	return CodeSuccess
}

func (s *server) Start() {
	s.router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
