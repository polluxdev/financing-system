package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/financing-system/config"
	"github.com/polluxdev/financing-system/helper"
	v1 "github.com/polluxdev/financing-system/internal/controller/http/v1"
	"github.com/polluxdev/financing-system/internal/repository"
	"github.com/polluxdev/financing-system/internal/service"
	"github.com/polluxdev/financing-system/pkg/httpserver"
	"github.com/polluxdev/financing-system/pkg/logger"
	"github.com/polluxdev/financing-system/pkg/postgres"
	"github.com/polluxdev/financing-system/pkg/validator"
)

func Run(config *config.Config) {
	log := logger.New(config.LogConfig.Level)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.PostgresConfig.Host,
		config.PostgresConfig.User,
		config.PostgresConfig.Password,
		config.PostgresConfig.Database,
		config.PostgresConfig.Port,
		config.PostgresConfig.SSLMode,
		config.PostgresConfig.TimeZone,
	)
	db, err := postgres.New(
		dsn,
		postgres.MaxPoolSize(config.PostgresConfig.MaxPoolSize),
		postgres.ConnAttempts(config.PostgresConfig.ConnAttempts),
		postgres.ConnTimeout(helper.GenerateTimeDuration(config.PostgresConfig.ConnTimeout, time.Second)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	validator := validator.New(log)
	validator.RegisterCustomValidator()

	tenorRepository := repository.NewTenorRepository()
	userFacilityDetailRepository := repository.NewUserFacilityDetailRepository()
	userFacilityLimitRepository := repository.NewUserFacilityLimitRepository()
	userFacilityRepository := repository.NewUserFacilityRepository()
	userRepository := repository.NewUserRepository()

	financeService := service.NewFinanceService(
		log,
		db,
		tenorRepository,
		userFacilityDetailRepository,
		userFacilityLimitRepository,
		userFacilityRepository,
	)
	tenorService := service.NewTenorService(log, db, tenorRepository)
	userService := service.NewUserService(log, db, userRepository)

	httpHandler := gin.New()
	v1.NewRouter(httpHandler, log, config, validator, financeService, tenorService, userService)
	httpServer := httpserver.New(httpHandler, httpserver.Port(config.HTTPConfig.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		log.Info("shutting down")
	case err := <-httpServer.Notify():
		log.Error(err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(err)
	}
}
