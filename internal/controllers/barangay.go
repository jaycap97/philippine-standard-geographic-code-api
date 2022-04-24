package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (s *Service) GetBarangayList(c *gin.Context) {
	bs, err := s.service.GetBarangayList()

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, bs)
}

func (s *Service) GetBarangay(c *gin.Context) {
	b, err := s.service.GetBarangay(c.Param("psgc_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(200, b)
}

func (s *Service) CreateBarangay(c *gin.Context) {
	var json domain.Barangay
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, err := s.service.CreateBarangay(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, b)
}

func (s *Service) UpdateBarangay(c *gin.Context) {
	var json domain.Barangay
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	b, err := s.service.UpdateBarangay(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, b)
}

func (s *Service) DeleteBarangay(c *gin.Context) {
	var json domain.Barangay
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	b, err := s.service.DeleteBarangay(&json)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, b)
}
