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
	s.Router.GET("", s.ListAllHandler)
	s.Router.GET("/notworn/:id", s.GetNotWornByIdHandler)
	s.Router.DELETE("/:id", s.DeleteNotWornHandler)
	s.Router.PATCH("/:id", s.UpdateNotWornHandler)
	s.Router.Static("/images", "./assets")
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

	file, _ := c.FormFile("image")
	err = s.DB.Model(&notworn).Update("file_name", file.Filename).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	// Upload the file to specific dst.
	fileName := fmt.Sprintf("%d_%s", notworn.ID, file.Filename)
	err = c.SaveUploadedFile(file, "./assets/"+fileName)
	if err != nil {
		// If any error occurs; we must delete created product
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})

		return
	}

	// Update image path of the product
	notworn.ImagePath = "/images/" + fileName
	err = s.DB.Save(&notworn).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})

		return
	}

	c.IndentedJSON(http.StatusCreated, &notworn)

}

func (s *Server) AddImageHandler(c *gin.Context) {

	var userObj file
	if err := c.ShouldBind(&userObj); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	obj, err := GetNotWorn(s.DB, intId)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	err = WriteFileName(s.DB, &obj, &userObj)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	stringID := fmt.Sprintf("%d", obj.ID)
	err = c.SaveUploadedFile(userObj.Avatar, "assets/"+stringID+userObj.Avatar.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   userObj,
	})

}

func (s *Server) ListAllHandler(c *gin.Context) {
	allNotWorn, err := ListAllNotWorn(s.DB)
	if err = c.ShouldBind(&allNotWorn); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.IndentedJSON(http.StatusAccepted, allNotWorn)

}

func (s *Server) GetNotWornByIdHandler(c *gin.Context) {

	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	notworn, err := GetNotWorn(s.DB, intId)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}
	c.IndentedJSON(http.StatusAccepted, &notworn)
}

func (s *Server) DeleteNotWornHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	err = HardDeleteNotWorn(s.DB, intId)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "deleted ",
	})

}

func (s *Server) UpdateNotWornHandler(c *gin.Context) {

	var notworn NotWorn
	err := c.ShouldBind(&notworn)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error", err)
		return
	}

	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	obj, err := UpdateNotWorn(s.DB, intId, &notworn)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, &obj)
}
