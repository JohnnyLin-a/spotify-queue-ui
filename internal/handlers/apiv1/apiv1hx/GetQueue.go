package apiv1hx

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
)

func GetQueue(runtimeContext *data.TRuntimeContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		queue, err := runtimeContext.SpotifyClient.GetQueue(ctx)

		if queue == nil || err != nil {
			log.Println(err)
			ctx.Data(http.StatusInternalServerError, "text/plain", []byte("Failed to load queue!"))
			return
		}
		partialGetQueue(queue).Render(ctx, ctx.Writer)
		ctx.AbortWithStatus(http.StatusOK)
	}
}
