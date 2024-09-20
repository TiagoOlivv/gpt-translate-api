package src

import (
	"gpt-translate-api/src/config"
	"gpt-translate-api/src/routes"
)

func Run() {
	config.LoadEnv()
	router := routes.SetupRouter()
	router.Run(":" + config.GetPort())
}
