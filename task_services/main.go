package main

import (
	"user-services/config"
	"user-services/docs"
	"user-services/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-chassis/go-chassis-extension/protocol/gin4r"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
	swaggerFiles "github.com/swaggo/files"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize logger" + err.Error())
	}
	defer logger.Sync()
}

// @title           User Manager API
// @version         1.0
// @description     This is a sample server for Task Manager.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5001
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.Title = "Task Manager API"
	docs.SwaggerInfo.Description = "This is a sample server for User Manager."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5001"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}
	gin4r.InstallPlugin()
	r := gin.Default()
	r.Use(cors.Default())
	dbConnector := config.ConnectDB()

	routes.UserRoutes(r, logger, dbConnector)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to User Service",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	chassis.RegisterSchema("rest", r)
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
