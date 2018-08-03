package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
	"log"
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username and password against DB after hashing the password
	if username == "foo" && password == "bar" {
		cookie := &http.Cookie{}

		// this is the same
		//cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		//TODO:  create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		//jwtToken to cookie
		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in!",
			"token": token,
		})
		//return c.String(http.StatusOK, "you were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong!")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
