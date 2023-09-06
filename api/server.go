package api

import (
	db "github/mh-hridoy/banking/db/sqlc"
	"net/http"

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
	router.GET("/account", server.GetListOfAccount)
	router.DELETE("/account/:id", server.DeleteSingleAccount)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "Route is not valid"}) })

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
