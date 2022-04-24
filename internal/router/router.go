package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/controllers"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/middleware"
)

var apiSecret = os.Getenv("API_SECRET_KEY")

type Server struct {
	api *gin.Engine
}

func Init(controller *controllers.Service) *Server {
	api := gin.New()
	api.Use(middleware.AuthAPIKey(apiSecret))
	api.GET("/regions", controller.GetRegionList)
	api.GET("/regions/:psgc_id", controller.GetRegion)
	api.POST("/regions", controller.CreateRegion)
	api.PUT("/regions", controller.UpdateRegion)
	api.DELETE("/regions", controller.DeleteRegion)

	api.GET("/provinces", controller.GetProvinceList)
	api.GET("/provinces/:psgc_id", controller.GetProvince)
	api.POST("/provinces", controller.CreateProvince)
	api.PUT("/provinces", controller.UpdateProvince)
	api.DELETE("/provinces", controller.DeleteProvince)

	api.GET("/cities", controller.GetCityList)
	api.GET("/cities/:psgc_id", controller.GetCity)
	api.POST("/cities", controller.CreateCity)
	api.PUT("/cities", controller.UpdateCity)
	api.DELETE("/cities", controller.DeleteCity)

	api.GET("/municipalities", controller.GetMunicipalityList)
	api.GET("/municipalities/:psgc_id", controller.GetMunicipality)
	api.POST("/municipalities", controller.CreateMunicipality)
	api.PUT("/municipalities", controller.UpdateMunicipality)
	api.DELETE("/municipalities", controller.DeleteMunicipality)

	api.GET("/sub-municipalities", controller.GetSubMunicipalityList)
	api.GET("/sub-municipalities/:psgc_id", controller.GetSubMunicipality)
	api.POST("/sub-municipalities", controller.CreateSubMunicipality)
	api.PUT("/sub-municipalities", controller.UpdateSubMunicipality)
	api.DELETE("/sub-municipalities", controller.DeleteSubMunicipality)

	api.GET("/barangays", controller.GetBarangayList)
	api.GET("/barangays/:psgc_id", controller.GetBarangay)
	api.POST("/barangays", controller.CreateBarangay)
	api.PUT("/barangays", controller.UpdateBarangay)
	api.DELETE("/barangays", controller.DeleteBarangay)

	return &Server{api}
}

func (s *Server) Start() {
	s.api.Run(":8080")
}
