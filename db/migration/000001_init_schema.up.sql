-- -- 创建数据库
-- CREATE DATABASE esight_database;

-- -- 连接到新创建的数据库
-- \c esight_database;

-- 创建设备信息表
-- CREATE TABLE device_info (
--     device_id SERIAL PRIMARY KEY,
--     device_name VARCHAR(255) NOT NULL,
--     device_type VARCHAR(50) NOT NULL,
--     ip_address VARCHAR(45) NOT NULL,
--     manufacturer VARCHAR(100),
--     model VARCHAR(100)
-- );
CREATE TABLE devices (
    device_id   SERIAL PRIMARY KEY,
    ip          VARCHAR(15) NOT NULL,
    hostname    VARCHAR(255) NOT NULL,
    model       VARCHAR(255) NOT NULL,
    vendor      VARCHAR(255) NOT NULL,
    device_type VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- 创建设备状态表
CREATE TABLE device_status (
    status_id SERIAL PRIMARY KEY,
    device_id INT NOT NULL,
    status_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL,
    cpu_usage FLOAT,
    memory_usage FLOAT,
    FOREIGN KEY (device_id) REFERENCES devices(device_id)
);

-- 创建用户信息表
CREATE TABLE user_info (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100),
    role VARCHAR(20) NOT NULL
);

-- 创建用户权限表
CREATE TABLE user_permission (
    permission_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    permission_type VARCHAR(50) NOT NULL,
    resource_id INT,
    FOREIGN KEY (user_id) REFERENCES user_info(user_id)
);

-- 创建拓扑关系表
CREATE TABLE topology_relationship (
    relationship_id SERIAL PRIMARY KEY,
    source_device_id INT NOT NULL,
    target_device_id INT NOT NULL,
    connection_type VARCHAR(50) NOT NULL,
    FOREIGN KEY (source_device_id) REFERENCES devices(device_id),
    FOREIGN KEY (target_device_id) REFERENCES devices(device_id)
);

-- 创建告警信息表
CREATE TABLE alarm_info (
    alarm_id SERIAL PRIMARY KEY,
    device_id INT NOT NULL,
    alarm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    alarm_level VARCHAR(20) NOT NULL,
    alarm_description TEXT,
    FOREIGN KEY (device_id) REFERENCES devices(device_id)
);