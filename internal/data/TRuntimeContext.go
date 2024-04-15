package data

import "github.com/zmb3/spotify/v2"

type TRuntimeContext struct {
	SelectedProfile   string
	AvailableProfiles []string
	SpotifyClient     *spotify.Client
}
