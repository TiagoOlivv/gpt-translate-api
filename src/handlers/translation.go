package handlers

import (
	"gpt-translate-api/src/models"
	"gpt-translate-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTranslation(c *gin.Context) {
	var request models.TranslationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	translatedText, err := services.TranslateTextWithChatGPT(request.Text, request.Source, request.Target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text":   request.Text,
		"source": request.Source,
		"target": request.Target,
		"result": translatedText,
	})
}
