package helpers

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func StartupSpotify(runtimeContext *data.TRuntimeContext, jsonPath string) error {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		runtimeContext.SpotifyClient = nil
		runtimeContext.SelectedProfile = ""
		return err
	}
	defer jsonFile.Close()
	var auth oauth2.Token
	b, err := io.ReadAll(jsonFile)
	if err != nil {
		runtimeContext.SpotifyClient = nil
		runtimeContext.SelectedProfile = ""
		return err
	}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		runtimeContext.SpotifyClient = nil
		runtimeContext.SelectedProfile = ""
		return err
	}
	spotifyClient := spotify.New(spotifyauth.New(
		spotifyauth.WithClientID(data.ConfigDatabase.SpotifyID),
		spotifyauth.WithClientSecret(data.ConfigDatabase.SpotifySecret),
	).Client(context.Background(), &auth))

	runtimeContext.SpotifyClient = spotifyClient

	return nil
}
