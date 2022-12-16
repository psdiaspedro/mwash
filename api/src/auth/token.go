package auth

import (
	"api/src/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GerarToken(usuarioID uint64, admin bool) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["exp"] = time.Now().Add(time.Hour * 12).Unix()
	permissoes["usuarioID"] = usuarioID
	permissoes["admin"] = admin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}