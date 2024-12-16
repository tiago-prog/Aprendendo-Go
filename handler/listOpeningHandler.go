package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSucess(ctx, "list-opening", openings)
}
