package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ndraaa15/musiku/internal/api/controller"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/domain/service"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
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
	db, err := mysql.NewMySqlClient()
	if err != nil {
		log.Printf("[musiku-server] failed to initialize musiku database : %v\n", err)
		return nil, err
	}
	log.Printf("[musiku-server] succes to initialize musiku database. Database connected\n")

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)

	s.router = gin.Default()
	s.ctrl = controller.NewController(us)

	mysql.Migration(db)

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()
	s.router.Run(fmt.Sprintf(":%s", os.Getenv("CONFIG_SERVER_PORT")))
	log.Printf("[musiku-server] Server is running at %s:%s", os.Getenv("CONFIG_SERVER_HOST"), os.Getenv("CONFIG_SERVER_PORT"))
	return CodeSuccess
}

func (s *server) Start() {
	log.Println("[musiku-server] starting server...")
	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, i'm musiku server"})
	})
}
