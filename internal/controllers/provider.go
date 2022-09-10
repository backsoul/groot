package controllers

// @title Google auth privder
// @version 1.0.0
// @host http://localhost:8000/api/sessions/oauth/google
// @BasePath /api
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/models"
	"github.com/backsoul/groot/pkg/services"
	"github.com/backsoul/groot/pkg/types"
	"github.com/backsoul/groot/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	uuid "github.com/nu7hatch/gouuid"
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
	u, _ := uuid.NewV4()
	uuid := fmt.Sprintf("%v", u)
	_, _ = models.CreateUser(uuid, user.Name, user.Email, "google", user.Picture)
	//TODO: valid if exist user
	// if err != nil {
	// 	return ctx.JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Error CreateUser",
	// 		"data":    err.Error(),
	// 	})
	// }
	User := types.User{}

	database.DB().Where("Email = ?", user.Email).First(&User)

	claims := types.UserClaims{
		Name:    User.Name,
		Email:   User.Email,
		Picture: User.Picture,
		Id:      User.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	type PayloadWallet struct {
		Payload interface{}
		User    types.UserClaims
	}
	var payloadWallet PayloadWallet
	var payloadEmpty interface{}
	payloadWallet.User = claims
	payloadWallet.Payload = payloadEmpty

	fmt.Println(payloadWallet)
	byts, _ := json.Marshal(payloadWallet)
	req, err := http.NewRequest("POST", "http://finance:8086/api/v1/wallet/create", bytes.NewBuffer(byts))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error microservice finance",
			"data":    err.Error(),
		})
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error client creating wallet",
			"data":    err.Error(),
		})
	}
	defer resp.Body.Close()

	defer resp.Body.Close()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtSecret := configs.Get("JWT_KEY")
	tokenJwt, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error Create JWT TOKEN",
			"data":    err.Error(),
		})
	}

	ctx.Cookie(services.AddNewCookie("access_token", tokenJwt, time.Now().Add(24*time.Hour)))
	url := configs.Get("REDIRECT_URL") + "?token=" + tokenJwt
	return ctx.Redirect(url, http.StatusTemporaryRedirect)
}
