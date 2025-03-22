package model

import (
	// "github.com/jackc/pgtype"
)


// DeviceRequest 定义post device的请求结构体，包含snmp请求
// type DeviceRequest struct {
// 	SNMPRequest SNMPRequest `json:"snmp_request"`
// 	IP          string      `json:"ip"`
// 	Hostname    string      `json:"hostname"`
// 	Model       string      `json:"model"`
// 	Vendor      string      `json:"vendor"`
// 	DeviceType  string      `json:"device_type"`
// }

// type DeviceResponse struct {
// 	DeviceDTO
// 	DeviceID  int32            `json:"device_id"`
// 	CreatedAt pgtype.Timestamp `json:"created_at"`
// 	UpdatedAt pgtype.Timestamp `json:"updated_at"`
// }

type DeviceIp struct{
	IP string `json:"ip"`
}