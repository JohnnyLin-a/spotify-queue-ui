package apiv1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/helpers"
	"github.com/johnnylin-a/spotify-queue-ui/internal/httppaths"
)

func ProfileSet(runtimeContext *data.TRuntimeContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		err := ctx.Request.ParseForm()
		if err != nil {
			log.Println("Error reading request body in apiv1.ProfileSet")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		profile := ctx.Request.PostFormValue("profile")
		if profile == "" {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := helpers.ProfileSet(runtimeContext, profile); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.Header("HX-Location", httppaths.SNEAK_SONGS)
		ctx.AbortWithStatus(http.StatusOK)
	}
}
