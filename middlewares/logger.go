package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(parametro gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			parametro.ClientIP,
			parametro.TimeStamp.Format(time.RFC822),
			parametro.Method,
			parametro.Path,
			parametro.StatusCode,
			parametro.Latency,
		)
	})
}
