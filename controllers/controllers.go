package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msajad79/pricing_engine/models"
)

func AddRulesController(ctx *gin.Context) {
	rule := []models.Rule{}
	fmt.Println(0)
	if err := ctx.ShouldBindJSON(&rule); err != nil {
		fmt.Println(1)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, rule)
	return
}
