-- 数据库初始化脚本
-- 此文件包含系统核心表的创建语句

-- 设备表 - 存储网络设备基本信息
CREATE TABLE "devices" (
    "device_id"   SERIAL PRIMARY KEY,
    "ip"          VARCHAR(15) NOT NULL,
    "hostname"    VARCHAR(255) NOT NULL,
    "model"       VARCHAR(255) NOT NULL,
    "vendor"      VARCHAR(255) NOT NULL,
    "device_type" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()), 
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

-- 用户表 - 系统用户账户
CREATE TABLE "users" ( 
    "id"          SERIAL PRIMARY KEY, 
    "username"    VARCHAR(64) NOT NULL UNIQUE, 
    "password"    VARCHAR(255) NOT NULL, 
    "role"        VARCHAR(255) NOT NULL,
    "active"      BOOLEAN DEFAULT true NOT NULL, 
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT (now()), 
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

-- 会话表 - 用户登录会话
CREATE TABLE "sessions" (
    "id"           uuid PRIMARY KEY,
    "username"     varchar(64) NOT NULL REFERENCES "users" ("username"),
    "refresh_token" varchar(64) NOT NULL,
    "user_agent"    varchar(64) NOT NULL,
    "client_ip"     varchar(64) NOT NULL,
    "is_blocked"    boolean NOT NULL DEFAULT false,
    "expires_at"    timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now())
);

-- 该表用于存储 SNMP 模板的基本信息
CREATE TABLE snmp_template (
    "id"          SERIAL PRIMARY KEY,
    "name"        VARCHAR(64) UNIQUE NOT NULL,
    "user_id"     INTEGER NOT NULL REFERENCES users("id"),
    "protocol"    VARCHAR(8) NOT NULL,
    "version"     VARCHAR(8) NOT NULL CHECK ("version" IN ('v2', 'v3')),
    "description" VARCHAR(255) NOT NULL,
    "device_count" INTEGER NOT NULL,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

-- SNMPv2参数表
CREATE TABLE snmpv2_parameters (
    "id"             SERIAL PRIMARY KEY,
    "template_id"    INTEGER NOT NULL REFERENCES snmp_template("id") ON DELETE CASCADE,
    "port"           INTEGER NOT NULL CHECK ("port" > 0 AND "port" <= 65535),
    
    -- 公共字段
    "timeout"        INTEGER NOT NULL CHECK ("timeout" > 0 AND "timeout" <= 60),
    "poll_interval"  INTEGER NOT NULL,
    "retries"        INTEGER NOT NULL CHECK ("retries" >= 0 AND "retries" <= 10),
    
    -- v2特有字段
    "read_community"  VARCHAR(64) NOT NULL,
    "write_community" VARCHAR(64) NOT NULL,
    "trap_community"  VARCHAR(64) NOT NULL
);

-- SNMPv3参数表
CREATE TABLE snmpv3_parameters (
    "id"            SERIAL PRIMARY KEY,
    "template_id"   INTEGER NOT NULL REFERENCES snmp_template("id") ON DELETE CASCADE,
    "port"          INTEGER NOT NULL CHECK ("port" > 0 AND "port" <= 65535),
    
    -- 公共字段
    "timeout"       INTEGER NOT NULL CHECK ("timeout" > 0 AND "timeout" <= 60),
    "poll_interval" INTEGER NOT NULL,
    "retries"       INTEGER NOT NULL CHECK ("retries" >= 0 AND "retries" <= 10),
    
    -- v3特有字段
    "security_level" VARCHAR(8) NOT NULL,
    "auth_protocol" VARCHAR(8) NOT NULL,
    "auth_password" VARCHAR(64) NOT NULL,
    "priv_protocol" VARCHAR(8) NOT NULL,
    "priv_password" VARCHAR(64) NOT NULL,
    "v3_user"       VARCHAR(64) NOT NULL,
    "engine_id"     VARCHAR(64) NOT NULL
);



-- 初始化管理员账户
INSERT INTO users (username, password, role)
VALUES ('admin', '$2a$10$GgZwysXaAfvE7ekE3dt8P.054MBCfFnHL/J/ikxfIxwihnqYSYgze', 'admin');
