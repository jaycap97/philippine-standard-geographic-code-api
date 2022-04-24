package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (s *Service) GetMunicipalityList(c *gin.Context) {
	ms, err := s.service.GetMunicipalityList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, ms)
}

func (s *Service) GetMunicipality(c *gin.Context) {
	m, err := s.service.GetMunicipality(c.Param("psgc_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(200, m)
}

func (s *Service) CreateMunicipality(c *gin.Context) {
	var json domain.Municipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	m, err := s.service.CreateMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, m)
}

func (s *Service) UpdateMunicipality(c *gin.Context) {
	var json domain.Municipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	m, err := s.service.UpdateMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, m)
}

func (s *Service) DeleteMunicipality(c *gin.Context) {
	var json domain.Municipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	m, err := s.service.DeleteMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, m)
}
