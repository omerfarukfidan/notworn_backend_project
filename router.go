package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	s.Router.PUT("/notworn/:id", s.AddImageHandler)
}

func (s *Server) CreateNotWornHandler(c *gin.Context) {

	var notworn NotWorn
	if err := c.ShouldBind(&notworn); err != nil {
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

func (s *Server) AddImageHandler(c *gin.Context) {

	var userObj file
	if err := c.ShouldBind(&userObj); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	id := c.Params.ByName("id")
	intVar, err := strconv.Atoi(id)
	obj, err := GetNotWorn(s.DB, intVar)
	fmt.Println(obj)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	WriteFileName(s.DB, &obj, &userObj)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	err = c.SaveUploadedFile(userObj.Avatar, "assets/"+userObj.Avatar.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   userObj,
	})

}
