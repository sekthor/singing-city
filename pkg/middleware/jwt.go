package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

	userId := c.Param("id")

	if userId == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	subject, err := token.Claims.GetSubject()

	if subject != userId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func validateToken(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte("aaa"), nil
}

func LoggedInAdmin(c *gin.Context) {

}
