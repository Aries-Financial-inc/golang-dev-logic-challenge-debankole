package routes

import (
	"ar/controllers"
	"ar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []model.OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := controllers.AnalyzeOptionsContracts(contracts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	return router
}
