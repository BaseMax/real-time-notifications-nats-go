package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var EXPTIME = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

func HashPassword(pass string) string {
	hashByte := sha256.Sum256([]byte(pass))
	hashStr := hex.EncodeToString(hashByte[:])
	return hashStr
}

func GetJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func CreateJwtToken(id uint, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(id),
		Issuer:    username,
		ExpiresAt: EXPTIME,
	})
	bearer, _ := token.SignedString(GetJwtSecret())
	return bearer
}

func GetLoggedinInfo(c echo.Context) (uint, string) {
	bearer := c.Request().Header.Get("Authorization")
	token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	username := claims["iss"].(string)
	id, _ := strconv.Atoi(claims["jti"].(string))
	return uint(id), username
}
