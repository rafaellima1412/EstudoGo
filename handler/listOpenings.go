package handler

import (
	"net/http"

	"example.com/estudoGo/config"
	"example.com/estudoGo/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath		/api/v1

//	@Summary		List Openings
//	@Description	List Openings jobs opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	ListOpeningsReponse
//	@Failure		500		{object}	ErrorReponse
//	@Router			/openings [get]

func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}
	db, err := config.InitializeMysql()
	if err != nil {
		logger.Errf("error connecting to the database: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error connecting to the database")
		return
	}
	// Se db for um ponteiro nulo, o erro ocorrerá ao chamar db.Create(&opening)
	if db == nil {
		logger.Err("db object is nil")
		sendError(c, http.StatusInternalServerError, "database object is nil")
		return
	}
	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusNotFound, "error listing openings")
		return
	}
	sendSuccess(c, "list-openings", openings)
}
