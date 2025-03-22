package apiserver

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/spf13/viper"
)

type ApiServer struct {
	Config *config.Config
	DBPool *pgxpool.Pool
}

func loadConfig() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 支持环境变量覆盖配置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("NDOT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config config.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &config, nil

}

func NewDatabasePool(dbConfig config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SslMode,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("解析数据库配置失败: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolConfig.MaxConns = int32(dbConfig.MaxOpenConns)
	poolConfig.MinConns = int32(dbConfig.MaxIdleConns)
	poolConfig.MaxConnLifetime = 90 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute
	poolConfig.ConnConfig.ConnectTimeout = 5 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("创建数据库连接池失败: %w", err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("测试数据库连接失败: %w", err)
	}

	return pool, nil
}

func NewApiServer() *ApiServer {

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}
	pool, err := NewDatabasePool(config.Database)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	app := &ApiServer{
		Config: config,
		DBPool: pool,
	}

	return app
}
