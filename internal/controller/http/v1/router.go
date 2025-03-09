package v1

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/financing-system/config"
	"github.com/polluxdev/financing-system/docs"

	"github.com/polluxdev/financing-system/internal/interfaces"
	"github.com/polluxdev/financing-system/internal/middleware"
	"github.com/polluxdev/financing-system/pkg/logger"
	"github.com/polluxdev/financing-system/pkg/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	handler *gin.Engine,
	logger logger.Interface,
	config *config.Config,
	validator validator.Validator,
	financeService interfaces.FinanceService,
	tenorService interfaces.TenorService,
	userService interfaces.UserService,
) {
	var basePath = "/api/v1"

	// Swagger
	docs.SwaggerInfo.Title = config.AppConfig.Name
	docs.SwaggerInfo.Version = config.AppConfig.Version
	docs.SwaggerInfo.Host = config.AppConfig.Host
	docs.SwaggerInfo.BasePath = basePath

	// Options
	handler.Use(gin.CustomRecovery(middleware.RecoverError))
	handler.Use(middleware.SetRequestID())

	// CORS
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Ping Route
	handler.GET("", func(c *gin.Context) {
		currentTime := time.Now().Format(time.RFC3339)
		serviceName := config.AppConfig.Name
		version := config.AppConfig.Version
		environment := config.AppConfig.Env

		response := gin.H{
			"requestId":   c.GetString("requestId"),
			"service":     serviceName,
			"message":     "Service is running smoothly!",
			"version":     version,
			"environment": environment,
			"timestamp":   currentTime,
			"status":      "success",
			"details": gin.H{
				"author":  "Pollux Dev Team",
				"support": "fahminur.dev@gmail.com",
			},
		}

		c.JSON(http.StatusOK, response)
	})

	// Healthz
	handler.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"requestId": c.GetString("requestId"),
			"status":    http.StatusOK,
			"message":   "OK",
		})
	})

	// 404
	handler.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"requestId": c.GetString("requestId"),
			"code":      http.StatusNotFound,
			"host":      c.Request.Host,
			"path":      c.Request.URL.Path,
			"method":    c.Request.Method,
			"message":   "Page not found",
		})
	})

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	group := handler.Group(basePath)
	{
		newFinanceRoutes(group, logger, config, validator, financeService)
		newTenorRoutes(group, logger, config, validator, tenorService)
		newUserRoutes(group, logger, config, validator, userService)
	}
}
