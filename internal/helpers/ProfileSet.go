package helpers

import (
	"errors"
	"slices"

	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
)

func ProfileSet(runtimeContext *data.TRuntimeContext, profile string) error {
	if !slices.Contains(runtimeContext.AvailableProfiles, profile) {
		return errors.New("error: profile not found")
	}
	runtimeContext.SelectedProfile = profile
	return nil
}
