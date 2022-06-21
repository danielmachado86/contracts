package api

import (
	"fmt"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/token"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.GET("/users/:username", server.getUser)

	authRoutes.POST("/contracts", server.createContract)
	authRoutes.GET("/contracts/:id", server.getContract)
	authRoutes.GET("/contracts", server.listContract)

	authRoutes.PUT("/contracts/:id/users/:username", server.createParty)
	authRoutes.GET("/contracts/:id/users/:username", server.getParty)
	authRoutes.GET("/contracts/:id/users", server.listParties)

	authRoutes.POST("/contracts/:id/periods", server.createPeriodParam)
	authRoutes.GET("/periods/:id", server.getPeriodParam)
	authRoutes.GET("/contracts/:id/periods", server.listPeriodParam)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
