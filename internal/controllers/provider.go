package controllers

// @title Google auth privder
// @version 1.0.0
// @host http://localhost:8000/api/sessions/oauth/google
// @BasePath /api
import (
	"net/http"
	"time"

	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/pkg/models"
	"github.com/backsoul/groot/pkg/services"
	"github.com/backsoul/groot/pkg/types"
	"github.com/backsoul/groot/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Code string `json:"code"`
}

// Provider Google
// @Summary provider google
// @Param token path string true "token auth provider"
// @Description auth session provider google
// @Success 307
// @Failure 404
// @Router /api/v1/sessions/oauth/google [get]
func ControllerAuthGoogleProvider(ctx *fiber.Ctx) error {
	var payload Payload
	payload.Code = ctx.Query("code")
	tokenRes, err := utils.GetGoogleOauthToken(payload.Code)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Errot get Google OAuth token",
			"data":    err.Error(),
		})
	}
	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error getting user",
			"data":    err.Error(),
		})
	}

	claims := types.UserClaims{
		Name:           user.Name,
		Email:          user.Email,
		Picture:        user.Picture,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtSecret := configs.Get("JWT_KEY")
	tokenJwt, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error SignedString",
			"data":    err.Error(),
		})
	}
	_, err = models.CreateUser(user.Name, user.Email, "google", user.Picture)
	ctx.Cookie(services.AddNewCookie("access_token", tokenJwt, time.Now().Add(24*time.Hour)))
	return ctx.Redirect("http://localhost:3000", http.StatusTemporaryRedirect)
}
