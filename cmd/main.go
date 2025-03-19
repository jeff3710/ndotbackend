package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeff3710/ndot/api"
	"github.com/jeff3710/ndot/config"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/pkg/log"
	"github.com/jeff3710/ndot/pkg/snmp"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config.Database:",log.String("err",err.Error()) )
	}
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.SSLMode,
	)
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return  
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return 
	}

	// 测试数据库连接
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("数据库连接测试失败:", log.String("err", err.Error()))
	}
	store:=db.NewStore(pool)
	snmp:= snmp.NewSNMPClient(config)
	runServer(store,*config,snmp)

}

func runServer(store db.Store, config config.Config,snmp *snmp.SNMPClient) {
	server:= api.NewServer(config, store,snmp)
	// if err!= nil {
	// 	log.Fatalf("创建服务器失败:", log.String("err", err.Error()))
	// }
// 将 int 类型的端口转换为 string 类型
portStr := fmt.Sprintf(":%d", config.App.Port)
err := server.Start(portStr)
	if err!= nil {
		log.Fatalf("服务器启动失败:", log.String("err", err.Error()))
	}
}
