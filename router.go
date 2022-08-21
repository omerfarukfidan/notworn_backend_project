package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	trustedProxies := []string{"127.0.0.1"}
	server := new(Server)
	server.Router = gin.Default()
	_ = server.Router.SetTrustedProxies(trustedProxies)
	server.DB = db

	return server

}

func (s *Server) InitRoutes() {
	s.Router.POST("", s.CreateNotWornHandler)

}

func (s *Server) CreateNotWornHandler(c *gin.Context) {

	var notworn NotWorn
	if err := c.BindJSON(&notworn); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})

		return
	}

	err := CreateNotWorn(s.DB, &notworn)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})

		return
	}

	c.IndentedJSON(http.StatusCreated, notworn)

}
