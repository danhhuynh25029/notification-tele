package server

import (
	"bot/config"
	"bot/pkg/mgo"
	"bot/service/domain/user/usecase"
	"bot/service/repository"

	"github.com/gin-gonic/gin"
)

type Domain struct {
	userUseCase usecase.IUserUsecase
}
type Server struct {
	router *gin.Engine
	cfg    config.Config
}

func NewServer(cfg config.Config) Server {
	router := gin.Default()

	return Server{
		cfg:    cfg,
		router: router,
	}
}
func (s *Server) Init() {
	// init mgo
	mgoClient := mgo.InitMgo(s.cfg)

	repo := repository.NewIRepo(mgoClient)
	s.initDomain(repo)
}

func (s *Server) initDomain(repo repository.IRepo) Domain {

	userUseCase := usecase.NewUserUseCase(repo)
	return Domain{
		userUseCase: userUseCase,
	}
}
