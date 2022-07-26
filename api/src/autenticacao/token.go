package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey)
}

func ValidarToken(r *http.Request) error {
	testarToken := extrairToken(r)
	if testarToken == "" {
		return errors.New("token invalido")
	}

	token, erro := jwt.Parse(testarToken, retornaChaveToken)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return nil
	}
	return errors.New("invalid permissions")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("authorization")
	if len(strings.Split(token, " ")) == 2 {

		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornaChaveToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method signing invalid %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtrairUsuarioIDToken(r *http.Request) (uint64, error) {

	testarToken := extrairToken(r)
	if testarToken == "" {
		return 0, errors.New("token invalido")
	}

	token, erro := jwt.Parse(testarToken, retornaChaveToken)

	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioId, nil

	}

	return 0, errors.New("token invalido")

}
