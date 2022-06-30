package api

import (
	"fmt"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/token"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	Logger     *zap.SugaredLogger
}

func NewServer() (*Server, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("cannot create logger: %w", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	server := &Server{
		Logger: sugar,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) ConfigServer(config utils.Config, store db.Store) error {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}

	server.tokenMaker = tokenMaker
	server.config = config
	server.store = store

	server.setupRouter()

	return nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("v1/users/login", server.loginUser)
	router.POST("v1/users", server.createUser)
	router.GET("v1/health", server.healthCheck)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.POST("v1/contracts", server.createContract)
	authRoutes.GET("v1/contracts/:id", server.getContract)
	authRoutes.GET("v1/contracts", server.listContract)

	authRoutes.POST("v1/contracts/:id/users", server.createParty)
	authRoutes.GET("v1/contracts/:id/users/:username", server.getParty)
	authRoutes.GET("v1/contracts/:id/users", server.listParties)

	authRoutes.POST("v1/contracts/:id/periods", server.createPeriodParam)
	authRoutes.GET("v1/periods/:id", server.getPeriodParam)
	authRoutes.GET("v1/contracts/:id/periods", server.listPeriodParam)

	server.router = router

}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
