package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (s *Service) GetCityList(c *gin.Context) {
	cts, err := s.service.GetCityList()

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, cts)
}

func (s *Service) GetCity(c *gin.Context) {
	ct, err := s.service.GetCity(c.Param("psgc_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(200, ct)
}

func (s *Service) CreateCity(c *gin.Context) {
	var json domain.City
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ct, err := s.service.CreateCity(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, ct)
}

func (s *Service) UpdateCity(c *gin.Context) {
	var json domain.City
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ct, err := s.service.UpdateCity(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, ct)
}

func (s *Service) DeleteCity(c *gin.Context) {
	var json domain.City
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ct, err := s.service.DeleteCity(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, ct)
}
