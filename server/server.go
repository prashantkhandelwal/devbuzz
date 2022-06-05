package server

import (
	"log"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/devbuzz/config"
	"github.com/prashantkhandelwal/devbuzz/server/handlers"
)

func Run() {

	c, err := config.InitConfig()
	if err != nil {
		log.Fatalf("ERROR: Cannot load configuration = %v", err)
	}

	port := c.Server.PORT

	if c.Environment != "" {
		if strings.ToLower(c.Environment) == "release" {
			gin.SetMode(gin.ReleaseMode)
		} else {
			gin.SetMode(gin.DebugMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	embedFS := EmbedFolder(Ui, "web", true)

	router.Use(static.Serve("/", embedFS))

	router.GET("/api/ping", handlers.Ping)
	//router.GET("/artist/search/:name", handlers.SearchArtist())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": "PAGE_NOT_FOUND", "message": "Page not found",
		})
	})

	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server! - %v", err)
	}

	log.Println("Server running!")

}
