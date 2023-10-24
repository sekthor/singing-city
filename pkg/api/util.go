package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *api) getUserIdFromContext(c *gin.Context) (int, error) {
	err := errors.New("could not get userid")
	val, ok := c.Get("userid")
	if !ok {
		return -1, err
	}

	userid, ok := val.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no userid"})
		return -1, err
	}

	return strconv.Atoi(userid)
}
