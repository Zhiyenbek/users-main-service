package handler

import (
	"errors"
	"log"
	"unicode"

	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *handler) SignIn(c *gin.Context) {
	req := &models.UserSignInRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}

	switch {
	case !validatePassword(req.Password):
		log.Println("invalid password")
		c.JSON(400, sendResponse(-1, nil, models.ErrInvalidPasswordFormat))
		return
	}
	var errMsg error
	tokens, err := h.service.AuthService.Login(req)
	if err != nil {
		log.Printf("Error occurred while login: %v", err)
		switch {
		case errors.Is(err, models.ErrWrongPassword):
			errMsg = models.ErrWrongPassword
		default:
			c.JSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
		log.Printf("Error occurred while login: %v", err)
		c.JSON(200, sendResponse(-1, nil, errMsg))
		return
	}
	c.SetCookie("access_token", tokens.AccessToken.TokenValue, int(tokens.AccessToken.ExpiresAt.Seconds()), "/", "", false, false)
	c.SetCookie("refresh_token", tokens.RefreshToken.TokenValue, int(tokens.RefreshToken.ExpiresAt.Seconds()), "/refresh-token", "", false, false)
	c.JSON(200, sendResponse(0, nil, nil))
}
func (h *handler) RefreshToken(c *gin.Context) {
	rtToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(401, sendResponse(-1, nil, models.ErrInvalidToken))
		return
	}
	tokens, err := h.service.AuthService.RefreshToken(rtToken)
	if err != nil {
		c.AbortWithStatusJSON(401, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	c.SetCookie("access_token", tokens.AccessToken.TokenValue, int(tokens.AccessToken.ExpiresAt), "/", "", false, false)
	c.SetCookie("refresh_token", tokens.RefreshToken.TokenValue, int(tokens.RefreshToken.ExpiresAt), "/refresh_token", "", false, false)
	c.JSON(200, sendResponse(0, nil, nil))
}

// validatePassword - function that validates password. Password being validated by these requirements:
// 1.Password length be between 8 - 16
// 2.Password must have upper case characters
// 3.Password must have special characters, except any type of quitation mark
// function returns boolean
func validatePassword(pass string) bool {
	if len(pass) < 8 || len(pass) > 16 {
		return false
	}
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range pass {
		switch {
		case char == 39 || char == 96 || char == 34: // ' ` " symbols in password
			return false
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
