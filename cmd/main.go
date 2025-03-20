package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/api/route"
	"github.com/jeff3710/ndot/pkg/log"
	"github.com/jeff3710/ndot/server"
)

func main() {
	app := server.App()
	config := app.Config
	options := &log.Options{
		DisableCaller:     config.Log.DisableCaller,
		DisableStacktrace: config.Log.DisableStacktrace,
		Level:             config.Log.Level,
		Format:            config.Log.Format,
		OutputPaths:       config.Log.OutputPaths,
	}

	log.Init(options)
	log.Infow("sugarlog config port", "端口", config.App.Port)
	defer log.Sync()

	pool := app.Pool
	defer pool.Close()

	gin := gin.Default()
	route.Setup(config, pool, gin)

	addr := fmt.Sprintf(":%d", config.App.Port)

	httpsrv := startInsecureServer(gin, addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("shutting down server at ", "addr", addr)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Fatalw("server forced to shutdown: %v", err)
	}

}

func startInsecureServer(g *gin.Engine, addr string) *http.Server {
	httpsrv := &http.Server{Addr: addr, Handler: g}

	log.Infow("start insecure server")
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw("http server listen: %s\n", err)
		}
	}()
	return httpsrv

}
