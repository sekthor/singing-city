package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogRequest(c *gin.Context) {
	start := time.Now()
	c.Next()
	stop := time.Now()
	latency := stop.Sub(start)

	var logEvent *zerolog.Event
	if c.Writer.Status() >= 500 {
		logEvent = log.Error()
	} else {
		logEvent = log.Info()
	}

	logEvent.
		Int("status", c.Writer.Status()).
		Str("client", c.ClientIP()).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("param", c.Request.URL.RawQuery). // http parameters
		Int("res_size", c.Writer.Size()).     // response size
		Str("latency", latency.String()).
		Msg(c.Errors.String())
}
