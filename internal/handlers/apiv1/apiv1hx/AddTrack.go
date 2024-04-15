package apiv1hx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/helpers"
)

func AddTrack(runtimeContext *data.TRuntimeContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		err := ctx.Request.ParseForm()
		if err != nil {
			ctx.Data(http.StatusInternalServerError, "text/plain", []byte("Failed to queue song!"))
			ctx.Abort()
			return
		}
		trackID := ctx.Request.PostFormValue("id")
		if trackID == "" {
			ctx.Data(http.StatusInternalServerError, "text/plain", []byte("Failed to queue song!"))
			ctx.Abort()
			return
		}

		err = helpers.QueueTrack(runtimeContext, trackID)
		if err != nil {
			ctx.Data(http.StatusInternalServerError, "text/plain", []byte("Failed to queue song!"))
			ctx.Abort()
			return
		}
		ctx.Header("HX-Trigger", "clear-textarea")
		ctx.Data(http.StatusOK, "text/plain", []byte("Added to queue!"))
		ctx.Abort()
	}
}
