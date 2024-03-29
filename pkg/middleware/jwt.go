package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

var ServerSecret []byte

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, validateToken)

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	subject, err := token.Claims.GetSubject()
	c.Set("userid", subject)
	c.Next()
}

// this can only be used when path is using :id as userID
func RequireResourceOwnerAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, validateToken)

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := c.Param("userid")

	if userId == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	subject, err := token.Claims.GetSubject()

	if subject != userId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("userid", subject)
	c.Next()
}

func validateToken(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return ServerSecret, nil
}

func LoggedInAdmin(c *gin.Context) {

}

func SetServerSecret(secret string) {
	ServerSecret = []byte(secret)

	if len(ServerSecret) == 0 {
		ServerSecret := make([]byte, 64)
		_, err := rand.Read(ServerSecret)
		if err != nil {
			log.Fatal().Err(err).Msg("could not generate server signing key")
		}
	}
}
