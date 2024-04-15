package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/httppaths"
)

func SelectedProfileMiddleware(runtimeContext *data.TRuntimeContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		if runtimeContext.SelectedProfile == "" {
			c.Header("HX-Location", httppaths.PROFILE_SELECT)
			c.Redirect(http.StatusTemporaryRedirect, httppaths.PROFILE_SELECT)
			c.Abort()
			return
		}
		c.Next()
	}
}
