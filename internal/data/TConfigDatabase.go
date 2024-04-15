package data

type TConfigDatabase struct {
	SpotifyID     string `env:"SPOTIFY_ID"`
	SpotifySecret string `env:"SPOTIFY_SECRET"`
	WebserverHost string `env:"WEBSERVER_HOST"`
}

var ConfigDatabase TConfigDatabase
