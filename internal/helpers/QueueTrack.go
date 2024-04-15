package helpers

import (
	"context"
	"errors"

	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/zmb3/spotify/v2"
)

func QueueTrack(runtimeContext *data.TRuntimeContext, trackID string) error {
	err := runtimeContext.SpotifyClient.QueueSong(context.Background(), spotify.ID(trackID))
	if err != nil {
		return errors.New("Failed to queue track " + trackID)
	}
	return nil
}
