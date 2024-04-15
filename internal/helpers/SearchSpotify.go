package helpers

import (
	"context"
	"strings"

	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/zmb3/spotify/v2"
)

func SearchSpotify(runtimeContext *data.TRuntimeContext, s string) ([]spotify.FullTrack, error) {
	// URL search
	if strings.HasPrefix(s, "spotify:track:") || strings.HasPrefix(s, "https://open.spotify.com/track/") {
		TrackID := ""
		if strings.HasPrefix(s, "https://open.spotify.com/track/") {
			TrackID = strings.TrimPrefix(s, "https://open.spotify.com/track/")
			if strings.Contains(TrackID, "?") {
				TrackID = TrackID[:strings.Index(TrackID, "?")]
			}
		} else {
			TrackID = strings.TrimPrefix(s, "spotify:track:")
		}
		result, err := runtimeContext.SpotifyClient.GetTrack(context.Background(), spotify.ID(TrackID))
		if err != nil {
			return []spotify.FullTrack{}, err
		}
		return []spotify.FullTrack{*result}, nil
	}
	// Text search
	ctx := context.Background()
	result, err := runtimeContext.SpotifyClient.Search(ctx, s, spotify.SearchTypeTrack, spotify.Limit(5))
	if err != nil {
		return []spotify.FullTrack{}, err
	}
	if len(result.Tracks.Tracks) > 0 {
		return result.Tracks.Tracks, nil
	} else {
		return []spotify.FullTrack{}, err
	}

}
