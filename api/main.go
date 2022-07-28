package main

import (
	"context"
	"encoding/gob"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/winnix/proreporting/api/config"
	"github.com/winnix/proreporting/api/controller"
	"github.com/winnix/proreporting/api/dbconn"
	"github.com/winnix/proreporting/api/middleware"
	"github.com/winnix/proreporting/api/web"
)

type options struct {
	config string
}

func initOptions() options {
	var opts options
	flag.StringVar(&opts.config, "config", "", "Path to config")
	flag.Parse()

	return opts
}

func main() {
	opts := initOptions()

	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	logger.Info("Args", os.Args)

	cfg, err := config.LoadConfig(opts.config)
	if err != nil {
		logger.WithField("configPath", opts.config).Fatal("Can't read config", err)
		return
	}

	if err := web.InitAuthProviders(cfg); err != nil {
		logger.Fatal("Can't initialize authorization provider(s)", err)
		return
	}

	dbClient, dbContext, dbCancel, err := dbconn.InitDB(cfg, logger)
	defer dbCancel()
	if dbClient != nil {
		defer dbClient.Disconnect(dbContext)
	}
	if err != nil {
		logger.Fatal("Can't connect to DB", err)
		return
	}

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: initRouter(cfg, logger),
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	gracefulShutdown(srv)
}

func initRouter(cfg *config.Config, logger *logrus.Logger) *gin.Engine {
	router := gin.New()

	gob.Register(web.CurrentUser{})
	web.InitSessions(router, cfg)

	router.Use(gin.Recovery())
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Validator())
	router.Use(middleware.Config(cfg))

	router.LoadHTMLGlob("templates/*")

	controller.SetRoutes(router, cfg)

	return router
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
