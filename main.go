package main

import (
	"fmt"
	"go-gin-api/docs"
	"go-gin-api/src/config"
	"go-gin-api/src/ioc"
	"go-gin-api/src/middleware"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	goServer "github.com/openscriptsin/go-server"
)

func main() {
	fmt.Println("go gin api")

	kernal := ioc.NewKernal()
	ginServer := goServer.New(
		kernal,
		ioc.RegisterControllers,
		middleware.CORSMiddleware(),
		middleware.GlobalIpBasedRateLimiterMiddleware(),
		middleware.ClsMiddleware(),
	)
	ginApp := ginServer.GetEngine()

	// registering swagger
	docs.SwaggerInfo.BasePath = "/"
	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Serve frontend static files
	ginApp.Use(static("client/dist"))

	ginServer.Start(config.GetConfig().Port)

}

// static serves static files from the build directory
func static(staticDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If the request is an API call, skip static file serving
		if len(c.Request.URL.Path) >= 4 && c.Request.URL.Path[:4] == "/api" {
			c.Next()
			return
		}
		//TODO: if path is not found, serve 404 page
		// Try to serve the file from the static directory
		path := filepath.Join(staticDir, c.Request.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// If the file doesn't exist, serve the index.html (for SPA routing)
			c.File(filepath.Join(staticDir, "index.html"))
			return
		}

		// Serve the file
		c.File(path)
		c.Abort()
	}
}
