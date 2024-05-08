package api

import (
	//"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"github.com/go-playground/validator/v10"
	db "FitnessProject/db/sqlc"
	//"FitnessProject/util"

)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}