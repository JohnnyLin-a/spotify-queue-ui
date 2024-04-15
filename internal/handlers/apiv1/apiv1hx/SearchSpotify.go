package apiv1hx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/helpers"
)

func SearchSpotify(runtimeContext *data.TRuntimeContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		q := ctx.Request.URL.Query().Get("q")
		if q == "" {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
		tracks, err := helpers.SearchSpotify(runtimeContext, q)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if len(tracks) == 1 {
			helpers.QueueTrack(runtimeContext, tracks[0].ID.String())
			ctx.Data(http.StatusOK, "text/plain", []byte("Added to queue!"))
			ctx.Abort()
			return
		}

		partialSearchSpotify(tracks).Render(ctx, ctx.Writer)
		ctx.AbortWithStatus(http.StatusOK)
	}
}
