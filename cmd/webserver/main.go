package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/spotify-queue-ui/internal/httppaths"
)

type tRuntimeContext struct {
	SelectedProfile   string
	AvailableProfiles []string
}

var runtimeContext = tRuntimeContext{
	SelectedProfile:   "",
	AvailableProfiles: []string{},
}

var SelectedProfileMiddleware gin.HandlerFunc = func(c *gin.Context) {
	if runtimeContext.SelectedProfile == "" {
		c.Header("HX-Location", httppaths.PROFILE_SELECT)
		c.Redirect(http.StatusTemporaryRedirect, httppaths.PROFILE_SELECT)
		c.Abort()
		return
	}
	c.Next()
}

func main() {
	// Load profiles
	entries, err := os.ReadDir("./profiles")
	if err != nil {
		log.Fatal("Did not find `profiles` folder!")
	}

	for _, e := range entries {
		if e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			runtimeContext.AvailableProfiles = append(runtimeContext.AvailableProfiles, e.Name())
		}
	}
	if len(runtimeContext.AvailableProfiles) == 1 {
		runtimeContext.SelectedProfile = runtimeContext.AvailableProfiles[0]
	}

	r := gin.Default()

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
		profileSelect().Render(c, c.Writer)
	})

	cp := r.Group(httppaths.SNEAK_SONGS)
	{
		cp.Use(SelectedProfileMiddleware)
		cp.GET(httppaths.ROOT, func(ctx *gin.Context) {
			sneakSongs().Render(ctx, ctx.Writer)
		})
	}

	log.Fatalln(r.Run("0.0.0.0:3002"))
}
