# NDOT 网络设备管理平台

## 项目概述
基于Go语言开发的网络设备管理平台，采用以下技术栈：
- 后端框架：Go 1.21 + Gin
- 数据库：PostgreSQL 15
- 容器化：Docker Compose
- 配置管理：YAML
- 接口文档：Swagger

## 快速启动
```bash
docker-compose up -d ndot-postgres  # 启动数据库
make run  # 启动应用
```

## 配置说明（config.yaml）
```yaml
app:
  port: 8081  # 服务监听端口

database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
```

## API文档
访问 [Swagger UI](http://localhost:8081/swagger/index.html) 查看接口文档

## SNMP配置
```yaml
system_oids:
  sysDescr: "1.3.6.1.2.1.1.1.0"
  sysObjectID: "1.3.6.1.2.1.1.2.0"
```

## 开发指南
1. 安装依赖：
```bash
go mod tidy
sqlc generate
```
2. 数据库迁移：
```bash
make migrate-up
```
3. 调试模式：
```bash
make debug
```

4. 临时hash密码

```
$2a$10$zI7jOObqZzS0yY1ia7lJiO5QHBrPZJ8y7I.IQhva0wh5g9E7VDrRi
```

