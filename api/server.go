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
	router.GET("/users/:username", server.getUser)

	router.POST("/contracts", server.createContract)
	router.GET("/contracts/:id", server.getContract)
	router.GET("/contracts", server.listContract)

	router.PUT("/contracts/:id/users/:username", server.createParty)
	router.GET("/contracts/:id/users/:username", server.getParty)
	router.GET("/contracts/:id/users", server.listParties)

	router.POST("/contracts/:id/periods", server.createPeriodParam)
	router.GET("/periods/:id", server.getPeriodParam)
	router.GET("/contracts/:id/periods", server.listPeriodParam)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
