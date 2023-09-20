package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"shorty/utils"

	"github.com/labstack/echo/v4"
)

func MapTokensController(app *echo.Echo) {
	tokenGroup := app.Group("/api/v1/tokens/")

	tokenGroup.GET("generate-csrf", generateCsrf)
	tokenGroup.GET("validate", decodeCsrf)
}

var key = []byte("THIS IS A SECURE KEY")

func generateCsrf(c echo.Context) error {
	// TODO: need refactoring
	buf := new(bytes.Buffer)
	buf.Write([]byte(utils.GenerateRandomHash(4)))

	hash := hmac.New(sha256.New, key)
	hash.Write(buf.Bytes())

	buf.Write(hash.Sum(nil))

	hashString := base64.StdEncoding.EncodeToString(buf.Bytes())
	return c.JSON(http.StatusOK, &echo.Map{
		"token": hashString,
	})
}

func decodeCsrf(c echo.Context) error {
	token := c.QueryParam("token")

	if len(token) != 48 {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"valid": false,
			"token": token,
		})
	}

	dest, err := base64.RawStdEncoding.DecodeString(token)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hash := hmac.New(sha256.New, key)
	fmt.Printf("LENGTH %d", len(dest))
	hash.Write(dest[:4])

	if !sliceEquals(hash.Sum(nil), dest[4:]) {

		return c.JSON(http.StatusForbidden, &echo.Map{
			"valid": false,
			"token": token,
		})
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"valid": true,
		"token": token,
	})
}

func sliceEquals(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
