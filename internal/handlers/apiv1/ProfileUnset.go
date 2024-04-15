package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/helpers"
	"github.com/johnnylin-a/spotify-queue-ui/internal/httppaths"
)

func ProfileUnset(runtimeContext *data.TRuntimeContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		helpers.ProfileUnset(runtimeContext)
		ctx.Header("HX-Location", httppaths.PROFILE_SELECT)
	}
}
