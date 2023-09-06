package api

import (
	db "github/mh-hridoy/banking/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.POST("/account", server.CreateAccount)
	router.GET("/account/:id", server.GetSingleAccount)

	server.route = router
	return server
}

func (s *Server) StartServer() error {
	err := s.route.Run()

	return err
}

func errorHandler(e error) gin.H {
	return gin.H{"error": e.Error()}
}
