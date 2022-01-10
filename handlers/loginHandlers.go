package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ustadho/belajar-go-auth/models"
	"github.com/ustadho/belajar-go-auth/token"
)

func LoginHandler(context *gin.Context) {
	var loginObj models.LoginRequest
	if err := context.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}

	// validate the loginObj for valid credential and if these are valid then

	var claims = &models.JwtClaims{}
	claims.CompanyID = "CompanyId"
	claims.Username = loginObj.UserName
	claims.Roles = []string{"admin", "supervisor"}
	claims.Audience = context.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(10) * time.Minute)
	tokenString, err := token.GenerateToken(claims, expirationTime)

	if err != nil {
		badRequest(context, http.StatusBadRequest, "error in generating token", []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
	}
	ok(context, http.StatusOK, "token created", tokenString)
}
