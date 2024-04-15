package helpers

import "github.com/johnnylin-a/spotify-queue-ui/internal/data"

func ProfileUnset(runtimeContext *data.TRuntimeContext) {
	runtimeContext.SelectedProfile = ""
}
