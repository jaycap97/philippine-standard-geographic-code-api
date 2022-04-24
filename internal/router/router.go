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
	psgc := api.Group("/psgc")
	{
		psgc.GET("/regions", controller.GetRegionList)
		psgc.GET("/regions/:psgc_id", controller.GetRegion)
		psgc.POST("/regions", controller.CreateRegion)
		psgc.PUT("/regions", controller.UpdateRegion)
		psgc.DELETE("/regions", controller.DeleteRegion)

		psgc.GET("/provinces", controller.GetProvinceList)
		psgc.GET("/provinces/:psgc_id", controller.GetProvince)
		psgc.POST("/provinces", controller.CreateProvince)
		psgc.PUT("/provinces", controller.UpdateProvince)
		psgc.DELETE("/provinces", controller.DeleteProvince)

		psgc.GET("/cities", controller.GetCityList)
		psgc.GET("/cities/:psgc_id", controller.GetCity)
		psgc.POST("/cities", controller.CreateCity)
		psgc.PUT("/cities", controller.UpdateCity)
		psgc.DELETE("/cities", controller.DeleteCity)

		psgc.GET("/municipalities", controller.GetMunicipalityList)
		psgc.GET("/municipalities/:psgc_id", controller.GetMunicipality)
		psgc.POST("/municipalities", controller.CreateMunicipality)
		psgc.PUT("/municipalities", controller.UpdateMunicipality)
		psgc.DELETE("/municipalities", controller.DeleteMunicipality)

		psgc.GET("/sub-municipalities", controller.GetSubMunicipalityList)
		psgc.GET("/sub-municipalities/:psgc_id", controller.GetSubMunicipality)
		psgc.POST("/sub-municipalities", controller.CreateSubMunicipality)
		psgc.PUT("/sub-municipalities", controller.UpdateSubMunicipality)
		psgc.DELETE("/sub-municipalities", controller.DeleteSubMunicipality)

		psgc.GET("/barangays", controller.GetBarangayList)
		psgc.GET("/barangays/:psgc_id", controller.GetBarangay)
		psgc.POST("/barangays", controller.CreateBarangay)
		psgc.PUT("/barangays", controller.UpdateBarangay)
		psgc.DELETE("/barangays", controller.DeleteBarangay)
	}

	return &Server{api}
}

func (s *Server) Start() {
	s.api.Run(":8080")
}
