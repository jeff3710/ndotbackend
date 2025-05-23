// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	DeviceID   int32     `json:"device_id"`
	Ip         string    `json:"ip"`
	Hostname   string    `json:"hostname"`
	Model      string    `json:"model"`
	Vendor     string    `json:"vendor"`
	DeviceType string    `json:"device_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type SnmpTemplate struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	UserID      int32     `json:"user_id"`
	Protocol    string    `json:"protocol"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	DeviceCount int32     `json:"device_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Snmpv2Parameter struct {
	ID             int32  `json:"id"`
	TemplateID     int32  `json:"template_id"`
	Port           int32  `json:"port"`
	Timeout        int32  `json:"timeout"`
	PollInterval   int32  `json:"poll_interval"`
	Retries        int32  `json:"retries"`
	ReadCommunity  string `json:"read_community"`
	WriteCommunity string `json:"write_community"`
	TrapCommunity  string `json:"trap_community"`
}

type Snmpv3Parameter struct {
	ID            int32  `json:"id"`
	TemplateID    int32  `json:"template_id"`
	Port          int32  `json:"port"`
	Timeout       int32  `json:"timeout"`
	PollInterval  int32  `json:"poll_interval"`
	Retries       int32  `json:"retries"`
	SecurityLevel string `json:"security_level"`
	AuthProtocol  string `json:"auth_protocol"`
	AuthPassword  string `json:"auth_password"`
	PrivProtocol  string `json:"priv_protocol"`
	PrivPassword  string `json:"priv_password"`
	V3User        string `json:"v3_user"`
	EngineID      string `json:"engine_id"`
}

type User struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
