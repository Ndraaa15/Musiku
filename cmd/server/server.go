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
	"github.com/Ndraaa15/musiku/internal/domain/entity"
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

	if err := mysql.Migration(db); err != nil {
		log.Printf("[musiku-server] failed to migrate musiku database : %v\n", err)
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	{
		var totalDays int64
		if err := db.Model(&entity.Day{}).Count(&totalDays).Error; err != nil {
			log.Printf("[musiku-server] failed to count total days : %v\n", err)
			return nil, err
		}

		if totalDays == 0 {
			if err := repository.SeedDays(db); err != nil {
				log.Printf("[musiku-server] failed to seed days : %v\n", err)
				return nil, err
			}
		}
	}

	{
		var totalVenue int64
		if err := db.Model(&entity.Venue{}).Count(&totalVenue).Error; err != nil {
			log.Printf("[musiku-server] failed to count total venue : %v\n", err)
			return nil, err
		}
		if totalVenue == 0 {
			if err := repository.SeedVenue(db); err != nil {
				log.Printf("[musiku-server] failed to seed venue : %v\n", err)
				return nil, err
			}
		}
	}

	venueRepository := repository.NewVenueRepository(db)
	venueService := service.NewVenueService(venueRepository)

	{
		var totalInstrument int64
		if err := db.Model(&entity.Instrument{}).Count(&totalInstrument).Error; err != nil {
			log.Printf("[musiku-server] failed to count total instrument : %v\n", err)
			return nil, err
		}

		if totalInstrument == 0 {
			if err := repository.SeedInstrument(db); err != nil {
				log.Printf("[musiku-server] failed to seed instrument : %v\n", err)
				return nil, err
			}
		}
	}

	instrumentRepository := repository.NewInstrumentRepository(db)
	instrumentService := service.NewInstrumentService(instrumentRepository)

	studioRepository := repository.NewStudioRepository(db)
	studioService := service.NewStudioService(studioRepository)

	s.handler = handler.NewHandler(userService, venueService, instrumentService, studioService)

	s.router = gin.Default()

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("CONFIG_SERVER_PORT")), s.router); err != nil {
		return ErrInternalServer
	}

	return CodeSuccess
}

func (s *server) Start() {
	log.Printf("[musiku-server] Server is running at %s:%s", os.Getenv("CONFIG_SERVER_HOST"), os.Getenv("CONFIG_SERVER_PORT"))
	log.Println("[musiku-server] starting server...")

	s.router.Use(middleware.CORS())

	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, i'm musiku server"})
	})

	route := s.router.Group("/api/v1")

	user := route.Group("/user")
	user.POST("/register", s.handler.Register)
	user.POST("/login", s.handler.Login)
	user.PATCH("/verify/:id", s.handler.VerifyAccount)

	user.Use(middleware.ValidateJWTToken())
	user.PATCH("/update", s.handler.UpdateUser)
	user.PATCH("/photo-profile", s.handler.UploadPhotoProfile)

	venue := route.Group("/venue")
	venue.Use(middleware.ValidateJWTToken())
	venue.GET("", s.handler.GetAllVenue)
	venue.GET("/:id", s.handler.GetVenueByID)
	venue.PATCH("/:id", s.handler.RentVenue)

	instrument := route.Group("/instrument")
	instrument.Use(middleware.ValidateJWTToken())
	instrument.GET("", s.handler.GetAllInstrument)
	instrument.GET("/:id", s.handler.GetInstrumentByID)
	instrument.PATCH("/:id", s.handler.RentInstrument)
	instrument.GET("/rent/:id-instrument/province", s.handler.GetProvince)
	instrument.GET("/rent/:id-instrument/city", s.handler.GetCity)
	instrument.GET("/rent/:id-instrument/cost", s.handler.GetCost)

	studio := route.Group("/studio")
	studio.Use(middleware.ValidateJWTToken())
	studio.GET("", s.handler.GetAllStudio)
	studio.GET("/:id", s.handler.GetStudioByID)
	studio.PATCH("/:id", s.handler.RentStudio)
}
