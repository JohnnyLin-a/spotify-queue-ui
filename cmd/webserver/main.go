package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/johnnylin-a/spotify-queue-ui/internal/data"
	"github.com/johnnylin-a/spotify-queue-ui/internal/handlers/apiv1"
	"github.com/johnnylin-a/spotify-queue-ui/internal/handlers/apiv1/apiv1hx"
	"github.com/johnnylin-a/spotify-queue-ui/internal/handlers/middlewares"
	"github.com/johnnylin-a/spotify-queue-ui/internal/helpers"
	"github.com/johnnylin-a/spotify-queue-ui/internal/httppaths"
)

var runtimeContext = &data.TRuntimeContext{
	SelectedProfile:   "",
	AvailableProfiles: []string{},
	SpotifyClient:     nil,
}

func init() {
	cleanenv.ReadEnv(&data.ConfigDatabase)
	cleanenv.ReadConfig(".env", &data.ConfigDatabase)
}

func main() {
	// Load profiles
	entries, err := os.ReadDir(data.PROFILES_PATH)
	if err != nil {
		log.Fatal("Did not find `profiles` folder!")
	}

	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			runtimeContext.AvailableProfiles = append(runtimeContext.AvailableProfiles, e.Name())
		}
	}
	if len(runtimeContext.AvailableProfiles) == 1 {
		runtimeContext.SelectedProfile = runtimeContext.AvailableProfiles[0]
		if err := helpers.StartupSpotify(runtimeContext, data.PROFILES_PATH+runtimeContext.SelectedProfile); err != nil {
			log.Fatalln(err)
		}
	}

	r := gin.Default()

	configureFrontend(r)
	configureAPI(r)

	if data.ConfigDatabase.WebserverHost == "" {
		data.ConfigDatabase.WebserverHost = "0.0.0.0:3000"
	}

	if runtime.GOOS == "windows" {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://"+strings.ReplaceAll(data.ConfigDatabase.WebserverHost, "0.0.0.0", "localhost")).Start()
	}
	log.Fatalln(r.Run(data.ConfigDatabase.WebserverHost))
}

func configureAPI(r *gin.Engine) {
	{
		apiV1 := r.Group(httppaths.API_V1_PREFIX)
		apiV1.POST(httppaths.API_V1_PROFILES_UNSET, apiv1.ProfileUnset(runtimeContext))
		apiV1.POST(httppaths.API_V1_PROFILES_SET, apiv1.ProfileSet(runtimeContext))
		{
			apiV1HX := apiV1.Group(httppaths.HX)
			apiV1HX.GET(httppaths.API_V1_Q, apiv1hx.SearchSpotify(runtimeContext))
			apiV1HX.GET(httppaths.API_V1_GET_QUEUE, apiv1hx.GetQueue(runtimeContext))
			apiV1HX.POST(httppaths.API_V1_ADD_TRACK, apiv1hx.AddTrack(runtimeContext))
		}
	}
}

func configureFrontend(r *gin.Engine) {
	r.GET(httppaths.ROOT, func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, httppaths.PROFILE_SELECT)
		c.Abort()
	})

	r.GET(httppaths.PROFILE_SELECT, func(c *gin.Context) {
		if runtimeContext.SelectedProfile != "" {
			c.Redirect(http.StatusTemporaryRedirect, httppaths.SNEAK_SONGS)
			c.Abort()
			return
		}
		profileSelect(runtimeContext).Render(c, c.Writer)
	})

	{
		ss := r.Group(httppaths.SNEAK_SONGS)
		ss.Use(middlewares.SelectedProfileMiddleware(runtimeContext))
		ss.GET(httppaths.ROOT, func(ctx *gin.Context) {
			sneakSongs(runtimeContext.SelectedProfile).Render(ctx, ctx.Writer)
		})
	}
}
