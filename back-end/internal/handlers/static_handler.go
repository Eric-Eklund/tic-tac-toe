package handlers

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var reactStatic embed.FS

// SetupStaticFileMiddleware configures static file serving middleware
func (h *Handler) SetupStaticFileMiddleware(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
}

// SetupSPAFallback configures the NoRoute handler for SPA support
func (h *Handler) SetupSPAFallback(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		if len(c.Request.URL.Path) >= 4 && c.Request.URL.Path[:4] != "/api" {
			c.FileFromFS("/", h.getFileSystem())
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})
}

// getFileSystem returns the embedded file system for static files
func (h *Handler) getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(reactStatic, "../frontend/build")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
