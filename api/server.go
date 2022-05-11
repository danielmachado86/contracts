package api

import (
	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/contracts", server.createContract)
	router.GET("/contracts/:id", server.getContract)
	router.GET("/contracts", server.listContract)

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	router.PUT("/contracts/:id/users/:userID", server.createParty)
	router.GET("/contracts/:id/users/:userID", server.getParty)
	router.GET("/contracts/:id/users", server.listParties)

	router.POST("/contracts/:id/periods", server.createPeriodParam)
	router.GET("/periods/:id", server.getPeriodParam)
	router.GET("/contracts/:id/periods", server.listPeriodParam)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
