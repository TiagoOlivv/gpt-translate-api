package models

type TranslationRequest struct {
	Text string `json:"q" binding:"required"`
	Source string `json:"source" binding:"required"`
	Target string `json:"target" binding:"required"`
}
