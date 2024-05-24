package infrastructure

import (
	"net/http"
	"os"
	"project-layout/pkg/shared/middleware"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// NewServer creates a new instance of the Gin Engine
func NewServer(db *gorm.DB, logger *logrus.Logger) *gin.Engine {
	if os.Getenv("STAGE_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginServer := gin.New()

	// middleware
	ginServer.Use(gin.Logger())
	ginServer.Use(gin.Recovery())
	ginServer.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowWebSockets:  true,
		AllowFiles:       true,
	}))
	ginServer.Use(middleware.RequestID())

	// health check
	var onStage string
	if strings.Contains(os.Getenv("RELEASE_IMAGE"), "dev") {
		onStage = "/dev"
	} else {
		onStage = "/"
	}
	ginServer.GET(onStage+"/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server_version": SERVER_VERSION,
		})
	})

	return ginServer
}
