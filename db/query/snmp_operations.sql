-- CreateSnmpTemplateBase creates a new SNMP template base record
-- and returns the newly created template ID
-- name: CreateSnmpTemplateBase :one
INSERT INTO snmp_template (
  name, user_id, protocol, version, device_count, description
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- CreateSnmpV2Template creates SNMP v2c specific parameters
-- name: CreateSnmpV2Template :exec
INSERT INTO snmpv2_parameters (
  template_id, port, read_community, write_community, trap_community, 
  timeout, poll_interval, retries
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
);

-- CreateSnmpV3Template creates SNMP v3 specific parameters
-- name: CreateSnmpV3Template :exec
INSERT INTO snmpv3_parameters (
  template_id, port, security_level, auth_protocol, auth_password, 
  priv_protocol, priv_password, v3_user, engine_id, 
  timeout, poll_interval, retries
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
);

-- name: UpdateSnmpTemplateBase :exec
UPDATE snmp_template
SET name = $1,
    protocol = $2,
    version = $3,
    description = $4,
    updated_at = now()
WHERE id = $5;

-- name: UpdateSnmpV2Parameters :exec
UPDATE snmpv2_parameters
SET port = $1,
    read_community = $2,
    write_community = $3,
    trap_community = $4,
    timeout = $5,
    poll_interval = $6,
    retries = $7
WHERE template_id = $8;

-- name: UpdateSnmpV3Parameters :exec
UPDATE snmpv3_parameters
SET port = $1,
    security_level = $2,
    auth_protocol = $3,
    auth_password = $4,
    priv_protocol = $5,
    priv_password = $6,
    v3_user = $7,
    engine_id = $8,
    timeout = $9,
    poll_interval = $10,
    retries = $11
WHERE template_id = $12;

-- name: DeleteSnmpV2Parameters :exec
DELETE FROM snmpv2_parameters
WHERE template_id = $1;

-- name: DeleteSnmpV3Parameters :exec
DELETE FROM snmpv3_parameters
WHERE template_id = $1;

-- name: DeleteSnmpTemplateBase :exec
DELETE FROM snmp_template
WHERE id = $1;

-- GetSnmpTemplateBase retrieves base SNMP template information
-- name: GetSnmpTemplateBase :one
SELECT 
  id, name, user_id, protocol, version, description,
  device_count, created_at, updated_at
FROM snmp_template
WHERE id = $1;

-- GetSnmpV2Parameters retrieves SNMP v2c specific parameters
-- name: GetSnmpV2Parameters :one
SELECT 
  id, port, read_community, write_community, trap_community,
  timeout, poll_interval, retries
FROM snmpv2_parameters
WHERE template_id = $1;

-- GetSnmpV3Parameters retrieves SNMP v3 specific parameters
-- name: GetSnmpV3Parameters :one
SELECT 
  id, port, security_level, auth_protocol, auth_password,
  priv_protocol, priv_password, v3_user, engine_id,
  timeout, poll_interval, retries
FROM snmpv3_parameters
WHERE template_id = $1;

-- ListSnmpTemplatesBase retrieves paginated list of base SNMP template information
-- name: ListSnmpTemplatesBase :many
SELECT 
  id, name, user_id, protocol, version, description,
  device_count, created_at, updated_at
FROM snmp_template
ORDER BY id
LIMIT $1
OFFSET $2;

-- ListSnmpV2Parameters retrieves SNMP v2c parameters for multiple templates
-- name: ListSnmpV2Parameters :many
SELECT 
  template_id, id, port, read_community, write_community, trap_community,
  timeout, poll_interval, retries
FROM snmpv2_parameters
WHERE template_id = ANY($1::int[]);

-- ListSnmpV3Parameters retrieves SNMP v3 parameters for multiple templates
-- name: ListSnmpV3Parameters :many
SELECT 
  template_id, id, port, security_level, auth_protocol, auth_password,
  priv_protocol, priv_password, v3_user, engine_id,
  timeout, poll_interval, retries
FROM snmpv3_parameters
WHERE template_id = ANY($1::int[]);