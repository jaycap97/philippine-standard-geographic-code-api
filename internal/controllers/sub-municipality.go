package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (s *Service) GetSubMunicipalityList(c *gin.Context) {
	sms, err := s.service.GetSubMunicipalityList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, sms)
}

func (s *Service) GetSubMunicipality(c *gin.Context) {
	sm, err := s.service.GetSubMunicipality(c.Param("psgc_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(200, sm)
}

func (s *Service) CreateSubMunicipality(c *gin.Context) {
	var json domain.SubMunicipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	sm, err := s.service.CreateSubMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, sm)
}

func (s *Service) UpdateSubMunicipality(c *gin.Context) {
	var json domain.SubMunicipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	sm, err := s.service.UpdateSubMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, sm)
}

func (s *Service) DeleteSubMunicipality(c *gin.Context) {
	var json domain.SubMunicipality
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	sm, err := s.service.DeleteSubMunicipality(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, sm)
}
