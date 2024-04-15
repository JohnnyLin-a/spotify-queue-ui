package data

type TConfigDatabase struct {
	SpotifyID     string `env:"SPOTIFY_ID"`
	SpotifySecret string `env:"SPOTIFY_SECRET"`
}

var ConfigDatabase TConfigDatabase
