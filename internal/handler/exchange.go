package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prerec/media-mine/internal/models"
	"github.com/prerec/media-mine/internal/usecase"
	"net/http"
)

var (
	InvalidJsonBodyResp = "invalid json body"
)

// @Summary		Exchange
// @Tags			exchange
// @Description	make an exchange
// @Accept			json
// @Produce		json
// @Param			input	body		models.Request	true	"exchange data"
// @Success		200		{object}	models.Response
// @Failure		400		{object}	models.ErrorResponse
// @Failure		404		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Failure		default	{object}	models.ErrorResponse
// @Router			/api/exchange [post]
func (h *Handler) exchange(c *gin.Context) {
	var req models.Request
	var res models.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: InvalidJsonBodyResp,
		})
		return
	}

	combinations, err := usecase.CombinationsFinder(req.Amount, req.Banknotes)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	res.Exchanges = combinations

	c.JSON(http.StatusOK, res)
}
