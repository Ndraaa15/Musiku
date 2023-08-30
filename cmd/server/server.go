package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ndraaa15/musiku/internal/api/handler"
	"github.com/Ndraaa15/musiku/internal/application/repository"
	"github.com/Ndraaa15/musiku/internal/application/service"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
	"github.com/Ndraaa15/musiku/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type server struct {
	router  *gin.Engine
	server  *http.Server
	handler *handler.Handler
}

func New() (*server, error) {
	s := &server{
		router: gin.Default(),
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	db, err := mysql.NewMySqlClient()
	if err != nil {
		log.Printf("[musiku-server] failed to initialize musiku database : %v\n", err)
		return nil, err
	}
	log.Printf("[musiku-server] succes to initialize musiku database. Database connected\n")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	s.handler = handler.NewHandler(userService)

	s.router = gin.Default()

	if err := mysql.Migration(db); err != nil {
		log.Printf("[musiku-server] failed to migrate musiku database : %v\n", err)
		return nil, err
	}

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()
	s.router.Run(fmt.Sprintf(":%s", os.Getenv("CONFIG_SERVER_PORT")))
	return CodeSuccess
}

func (s *server) Start() {
	log.Printf("[musiku-server] Server is running at %s:%s", os.Getenv("CONFIG_SERVER_HOST"), os.Getenv("CONFIG_SERVER_PORT"))
	log.Println("[musiku-server] starting server...")

	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, i'm musiku server"})
	})

	route := s.router.Group("/api/v1")

	route.POST("/register", s.handler.Register)
	route.POST("/login", s.handler.Login)

	route.Use(middleware.ValidateJWTToken())
}
