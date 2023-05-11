package middleware

import (
	"shop/config"
	"shop/model"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// LoginOnly protect routes
func LoginOnly() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Get("SECRET")),
		ErrorHandler: JwtError,
	})
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Token")
		token, err := ParseToken(tokenStr)
		if err != nil {
			return JwtError(c, err)
		}

		if !token.IsAdmin {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON("Not Admin")
		}

		return c.Next()
	}
}

func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil, "error": err.Error()})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil, "error": err.Error()})
}

func GenerateJwtToken(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Nama
	claims["user_id"] = user.ID
	claims["toko_id"] = user.Toko.ID
	claims["is_admin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(config.Get("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type TokenData struct {
	Username string    `json:"username"`
	UserId   uint      `json:"user_id"`
	IsAdmin  bool      `json:"is_admin"`
	TokoID   uint      `json:"toko_id"`
	Expired  time.Time `json:"exp"`
}

func ParseToken(tokenString string) (TokenData, error) {
	secret := config.Get("SECRET")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return TokenData{}, err
	}

	username := claims["username"].(string)
	userid := uint(claims["user_id"].(float64))
	isAdmin := bool(claims["is_admin"].(bool))
	tokoId := uint(claims["toko_id"].(float64))

	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	return TokenData{
		Username: username,
		UserId:   userid,
		IsAdmin:  isAdmin,
		TokoID:   tokoId,
		Expired:  exp,
	}, nil
}
