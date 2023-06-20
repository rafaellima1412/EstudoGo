package handler

import (
	"net/http"

	"example.com/estudoGo/config"
	"example.com/estudoGo/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath		/api/v1

//	@Summary		Update Opening
//	@Description	Update Opening job
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id query string			true	"Opening indetification"
//	@Param			opening	body		UpdateOpeningRequest	true	"Opening data to update"
//	@Success		200		{object}	UpdateOpeningReponse
//	@Failure		400		{object}	ErrorReponse
//	@Failure		404		{object}	ErrorReponse
//	@Failure		500		{object}	ErrorReponse
//	@Router			/opening [put]
func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validade(); err != nil {
		logger.Errf("error validation: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}
	opening := schemas.Opening{}

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
	if err := db.First(&opening).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary < 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error updating opening on database")
		return
	}
	sendSuccess(c, "updating-opening", opening)

}
