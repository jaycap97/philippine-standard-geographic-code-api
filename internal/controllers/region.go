package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (s *Service) GetRegionList(c *gin.Context) {
	rs, err := s.service.GetRegionList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, rs)
}

func (s *Service) GetRegion(c *gin.Context) {
	r, err := s.service.GetRegion(c.Param("psgc_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(200, r)
}

func (s *Service) CreateRegion(c *gin.Context) {
	var json domain.Region
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	r, err := s.service.CreateRegion(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, r)
}

func (s *Service) UpdateRegion(c *gin.Context) {
	var json domain.Region
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	r, err := s.service.UpdateRegion(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, r)
}

func (s *Service) DeleteRegion(c *gin.Context) {
	var json domain.Region
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	r, err := s.service.DeleteRegion(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, r)
}
