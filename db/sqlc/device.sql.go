// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: device.sql

package db

import (
	"context"
)

const createDevice = `-- name: CreateDevice :exec
INSERT INTO devices (ip, hostname, model, vendor, device_type)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (ip) DO UPDATE
SET hostname = EXCLUDED.hostname,
    model = EXCLUDED.model,
    vendor = EXCLUDED.vendor,
    device_type = EXCLUDED.device_type
`

type CreateDeviceParams struct {
	Ip         string `json:"ip"`
	Hostname   string `json:"hostname"`
	Model      string `json:"model"`
	Vendor     string `json:"vendor"`
	DeviceType string `json:"device_type"`
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) error {
	_, err := q.db.Exec(ctx, createDevice,
		arg.Ip,
		arg.Hostname,
		arg.Model,
		arg.Vendor,
		arg.DeviceType,
	)
	return err
}

const deleteDevice = `-- name: DeleteDevice :exec
DELETE FROM devices
WHERE device_id = $1
`

func (q *Queries) DeleteDevice(ctx context.Context, deviceID int32) error {
	_, err := q.db.Exec(ctx, deleteDevice, deviceID)
	return err
}

const getDeviceById = `-- name: GetDeviceById :one
SELECT device_id, ip, hostname, model, vendor, device_type, created_at, updated_at FROM devices
WHERE device_id = $1
`

func (q *Queries) GetDeviceById(ctx context.Context, deviceID int32) (Device, error) {
	row := q.db.QueryRow(ctx, getDeviceById, deviceID)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.Ip,
		&i.Hostname,
		&i.Model,
		&i.Vendor,
		&i.DeviceType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDeviceByIp = `-- name: GetDeviceByIp :one
SELECT device_id, ip, hostname, model, vendor, device_type, created_at, updated_at FROM devices
WHERE ip = $1
`

func (q *Queries) GetDeviceByIp(ctx context.Context, ip string) (Device, error) {
	row := q.db.QueryRow(ctx, getDeviceByIp, ip)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.Ip,
		&i.Hostname,
		&i.Model,
		&i.Vendor,
		&i.DeviceType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDevices = `-- name: ListDevices :many
SELECT device_id, ip, hostname, model, vendor, device_type, created_at, updated_at FROM devices
ORDER BY device_id
LIMIT $1
OFFSET $2
`

type ListDevicesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDevices(ctx context.Context, arg ListDevicesParams) ([]Device, error) {
	rows, err := q.db.Query(ctx, listDevices, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Device{}
	for rows.Next() {
		var i Device
		if err := rows.Scan(
			&i.DeviceID,
			&i.Ip,
			&i.Hostname,
			&i.Model,
			&i.Vendor,
			&i.DeviceType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDeviceIP = `-- name: UpdateDeviceIP :exec
UPDATE devices
SET ip = $1
WHERE device_id = $2
`

type UpdateDeviceIPParams struct {
	Ip       string `json:"ip"`
	DeviceID int32  `json:"device_id"`
}

func (q *Queries) UpdateDeviceIP(ctx context.Context, arg UpdateDeviceIPParams) error {
	_, err := q.db.Exec(ctx, updateDeviceIP, arg.Ip, arg.DeviceID)
	return err
}

const updateDeviceManufacturer = `-- name: UpdateDeviceManufacturer :exec
UPDATE devices
SET vendor = $1
WHERE device_id = $2
`

type UpdateDeviceManufacturerParams struct {
	Vendor   string `json:"vendor"`
	DeviceID int32  `json:"device_id"`
}

func (q *Queries) UpdateDeviceManufacturer(ctx context.Context, arg UpdateDeviceManufacturerParams) error {
	_, err := q.db.Exec(ctx, updateDeviceManufacturer, arg.Vendor, arg.DeviceID)
	return err
}

const updateDeviceModel = `-- name: UpdateDeviceModel :exec
UPDATE devices
SET model = $1
WHERE device_id = $2
`

type UpdateDeviceModelParams struct {
	Model    string `json:"model"`
	DeviceID int32  `json:"device_id"`
}

func (q *Queries) UpdateDeviceModel(ctx context.Context, arg UpdateDeviceModelParams) error {
	_, err := q.db.Exec(ctx, updateDeviceModel, arg.Model, arg.DeviceID)
	return err
}

const updateDeviceName = `-- name: UpdateDeviceName :exec
UPDATE devices
SET hostname = $1
WHERE device_id = $2
`

type UpdateDeviceNameParams struct {
	Hostname string `json:"hostname"`
	DeviceID int32  `json:"device_id"`
}

func (q *Queries) UpdateDeviceName(ctx context.Context, arg UpdateDeviceNameParams) error {
	_, err := q.db.Exec(ctx, updateDeviceName, arg.Hostname, arg.DeviceID)
	return err
}

const updateDeviceType = `-- name: UpdateDeviceType :exec
UPDATE devices
SET device_type = $1
WHERE device_id = $2
`

type UpdateDeviceTypeParams struct {
	DeviceType string `json:"device_type"`
	DeviceID   int32  `json:"device_id"`
}

func (q *Queries) UpdateDeviceType(ctx context.Context, arg UpdateDeviceTypeParams) error {
	_, err := q.db.Exec(ctx, updateDeviceType, arg.DeviceType, arg.DeviceID)
	return err
}
