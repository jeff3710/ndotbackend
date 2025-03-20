package model

import (
	"github.com/jackc/pgtype"
)



type DeviceRequest struct {
	SNMPRequest SNMPRequest `json:"snmp_request"`
	IP          string      `json:"ip"`
	Hostname    string      `json:"hostname"`
	Model       string      `json:"model"`
	Vendor      string      `json:"vendor"`
	DeviceType  string      `json:"device_type"`
}

type DeviceResponse struct {
	DeviceDTO
	DeviceID  int32            `json:"device_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
