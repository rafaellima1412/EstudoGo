package handler

import (
	"net/http"

	"example.com/estudoGo/config"
	"example.com/estudoGo/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath		/api/v1

//	@Summary		Create Opening
//	@Description	Create  a new Opening job
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateOpeningRequest	true	"Request body"
//	@Success		200		{object}	CreateOpeningReponse
//	@Failure		400		{object}	ErrorReponse
//	@Failure		404		{object}	ErrorReponse
//	@Failure		500		{object}	ErrorReponse
//	@Router			/opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validade(); err != nil {
		logger.Errf("error validation: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	db, err := config.InitializeMysql()
	if err != nil {
		logger.Errf("error connecting to the database: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error connecting to the database")
		return
	}
	// Se db for um ponteiro nulo, o erro ocorrerá ao chamar db.Create(&opening)
	if db == nil {
		logger.Err("db object is nil")
		sendError(ctx, http.StatusInternalServerError, "database object is nil")
		return
	}
	//logger.Infof("request received: %+v", request)
	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
