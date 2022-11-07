package server

import (
	"bot/config"
	"bot/pkg/mgo"
	"bot/service/domain/user/delivery/http"
	"bot/service/domain/user/usecase"
	"bot/service/repository"

	httpSer "net/http"

	"github.com/gin-gonic/gin"
)

type Domain struct {
	userUseCase usecase.IUserUsecase
}
type Server struct {
	router     *gin.Engine
	cfg        config.Config
	HttpServer *httpSer.Server
}

func NewServer(cfg *config.Config) (Server, error) {
	router := gin.Default()

	return Server{
		cfg:    *cfg,
		router: router,
	}, nil
}
func (s *Server) Init() {
	// init mgo
	mgoClient := mgo.InitMgo(s.cfg)

	repo := repository.NewIRepo(mgoClient)
	domain := s.initDomain(repo)
	s.initCors()
	s.initRouter(domain)
}

func (s *Server) initDomain(repo repository.IRepo) Domain {

	userUseCase := usecase.NewUserUseCase(repo)
	return Domain{
		userUseCase: userUseCase,
	}
}
func (s *Server) initCors() {
}
func (s *Server) initRouter(domain Domain) {
	userHandle := http.NewUserHandler(domain.userUseCase)

	router := s.router.Group("v1")
	userHandle.NewUserRoute(router)
}
func (s *Server) StartServer() error {
	s.httpServer = &httpSer.Server{
		Handler: s.router,
		Addr:    ":8080",
	}
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
