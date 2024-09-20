package routes

import (
	"gpt-translate-api/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/translate", handlers.HandleTranslation)

	return router
}
