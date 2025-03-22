-- name: CreateDevice :exec
INSERT INTO devices (ip, hostname, model, vendor, device_type)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (ip) DO UPDATE
SET hostname = EXCLUDED.hostname,
    model = EXCLUDED.model,
    vendor = EXCLUDED.vendor,
    device_type = EXCLUDED.device_type;

-- name: UpdateDeviceAll :exec
UPDATE devices
SET hostname = $1,
    model = $2,
    vendor = $3,
    device_type = $4
WHERE ip = $5;

-- name: UpdateDeviceName :exec
UPDATE devices
SET hostname = $1
WHERE device_id = $2;

-- name: UpdateDeviceType :exec
UPDATE devices
SET device_type = $1
WHERE device_id = $2;

-- name: UpdateDeviceIP :exec
UPDATE devices
SET ip = $1
WHERE device_id = $2;

-- name: UpdateDeviceVendor :exec
UPDATE devices
SET vendor = $1
WHERE device_id = $2;

-- name: UpdateDeviceModel :exec
UPDATE devices
SET model = $1
WHERE device_id = $2;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE device_id = $1;

-- name: GetDeviceById :one
SELECT * FROM devices
WHERE device_id = $1;

-- name: GetDeviceByIp :one
SELECT * FROM devices
WHERE ip = $1;

-- name: ListDevices :many
SELECT * FROM devices
ORDER BY device_id
LIMIT $1
OFFSET $2;