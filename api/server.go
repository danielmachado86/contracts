package api

import (
	"fmt"

	db "github.com/danielmachado86/contracts/db"
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

	rVersion := router.Group(fmt.Sprintf("/%s", server.config.ApiVersion))

	// Authentication endpoints
	rVersion.POST("/sessions", server.createSessions)
	rVersion.DELETE("/sessions", server.deleteSessions)

	rVersion.POST("/users", server.createUser)
	rVersion.GET("/health", server.healthCheck)

	authRoutes := rVersion.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.POST("/contracts", server.createContract)
	authRoutes.GET("/contracts/:id", server.getContract)
	authRoutes.GET("/contracts", server.listContract)

	authRoutes.POST("/contracts/:id/users", server.createParty)
	authRoutes.GET("/contracts/:id/users/:username", server.getParty)
	authRoutes.GET("/contracts/:id/users", server.listParties)

	authRoutes.POST("/contracts/:id/signatures", server.createSignature)
	authRoutes.GET("/contracts/:id/signatures", server.listSignatures)

	server.router = router

}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}

func errorResponse(err error, code int) gin.H {
	return gin.H{
		"message": err.Error(),
		"code":    code,
	}
}
