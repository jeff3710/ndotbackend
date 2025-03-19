// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AlarmInfo struct {
	AlarmID          int32            `json:"alarm_id"`
	DeviceID         int32            `json:"device_id"`
	AlarmTime        pgtype.Timestamp `json:"alarm_time"`
	AlarmLevel       string           `json:"alarm_level"`
	AlarmDescription pgtype.Text      `json:"alarm_description"`
}

type Device struct {
	DeviceID   int32            `json:"device_id"`
	Ip         string           `json:"ip"`
	Hostname   string           `json:"hostname"`
	Model      string           `json:"model"`
	Vendor     string           `json:"vendor"`
	DeviceType string           `json:"device_type"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

type DeviceStatus struct {
	StatusID    int32            `json:"status_id"`
	DeviceID    int32            `json:"device_id"`
	StatusTime  pgtype.Timestamp `json:"status_time"`
	Status      string           `json:"status"`
	CpuUsage    pgtype.Float8    `json:"cpu_usage"`
	MemoryUsage pgtype.Float8    `json:"memory_usage"`
}

type TopologyRelationship struct {
	RelationshipID int32  `json:"relationship_id"`
	SourceDeviceID int32  `json:"source_device_id"`
	TargetDeviceID int32  `json:"target_device_id"`
	ConnectionType string `json:"connection_type"`
}

type UserInfo struct {
	UserID   int32       `json:"user_id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Email    pgtype.Text `json:"email"`
	Role     string      `json:"role"`
}

type UserPermission struct {
	PermissionID   int32       `json:"permission_id"`
	UserID         int32       `json:"user_id"`
	PermissionType string      `json:"permission_type"`
	ResourceID     pgtype.Int4 `json:"resource_id"`
}
