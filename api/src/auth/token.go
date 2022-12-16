package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func pegaToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func verificaAssinaturaToken(token *jwt.Token) (interface{}, error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Assinatura do token indefinida! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func PegaUsuarioIDToken(r *http.Request) (uint64, error) {
	tokenString := pegaToken(r)
	token, erro := jwt.Parse(tokenString, verificaAssinaturaToken)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token invalido")
}

func ValidaToken(r *http.Request) error {
	tokenString := pegaToken(r)
	token, erro := jwt.Parse(tokenString, verificaAssinaturaToken)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}