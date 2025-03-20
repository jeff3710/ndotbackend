package server

import (
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Config *Config
	Pool *pgxpool.Pool
}

func App() Application{
 app:= &Application{}
 var err error
 app.Config,err= LoadConfig()
 if err!= nil {
	log.Fatal("配置加载失败:", err)
 }
 app.Pool,err = NewDatabasePool(app.Config.Database)
 if err!= nil {
	log.Fatal("数据库连接失败:", err)
 }
 return *app
}