package main

import (
	"backend/internal/config"
	"backend/internal/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/test", controllers.Test)
	r.GET("/getRecord/:id", controllers.GetUniversityData)
	r.GET("/getAllRecords", controllers.GetUniversitiesData)
	r.GET("/getAllRecordsFilter", controllers.GetUniversitiesDataFilter)
	r.GET("/reportPlaces", controllers.GetRatingREPORT)
	r.GET("/reportFilters", controllers.GetReport)
	r.GET("/places", controllers.GetRatingJSON)
	r.PUT("/updateRecord", controllers.UpdateUniversity)
	r.PUT("/updatePlace", controllers.ChangePlace)
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.ProjectConfig.Deploy.Port))
}
