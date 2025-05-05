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
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/apiserver"
	"github.com/jeff3710/ndot/internal/apiserver/route"
	mw "github.com/jeff3710/ndot/internal/pkg/middleware"
	"github.com/jeff3710/ndot/pkg/log"
)

func main() {

	// 创建一个新的api server实例
	app, err := apiserver.NewApiServer()
	if err != nil {
		log.Fatalw("failed to create api server: %v", err)
	}

	config := app.Config

	options := &log.Options{
		DisableCaller:     config.Log.DisableCaller,
		DisableStacktrace: config.Log.DisableStacktrace,
		Level:             config.Log.Level,
		Format:            config.Log.Format,
		OutputPaths:       config.Log.OutputPaths,
	}
	token := app.TokenMaker
	// 初始化日志模块
	log.Init(options)
	defer log.Sync()

	pool := app.DBPool
	defer pool.Close()

	store := db.NewStore(pool)
	// 打印数据库连接信息

	g := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), mw.Cors}
	g.Use(mws...)

	// 注册路由
	route.Setup(config, store, token, g)

	addr := fmt.Sprintf(":%d", config.App.Port)

	httpsrv := startHttpServer(g, addr)
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

func startHttpServer(g *gin.Engine, addr string) *http.Server {
	httpsrv := &http.Server{Addr: addr, Handler: g}

	log.Infow("start http server")
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw("http server listen: %s\n", err)
		}
	}()
	return httpsrv

}
