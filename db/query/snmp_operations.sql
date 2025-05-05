-- name: CreateSnmpTemplateBase :one
INSERT INTO snmp_template (
  name, user_id, protocol, version, device_count, description
)VALUES($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: CreateSnmpTemplateWithV2Parameters :exec
INSERT INTO snmpv2_parameters (
  template_id, port, read_community, write_community, trap_community, timeout, poll_interval, retries
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
);

-- name: CreateSnmpTemplateWithV3Parameters :exec
INSERT INTO snmpv3_parameters (
  template_id, port, security_level, auth_protocol, auth_password, priv_protocol, priv_password, v3_user, engine_id, timeout, poll_interval, retries
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
);

-- name: UpdateSnmpTemplateWithParameters :exec
UPDATE snmp_template
SET name = $1,
    protocol = $2,
    version = $3,
    description = $4,
    updated_at = now()
WHERE id = $5;

UPDATE snmpv2_parameters
SET port = $6,
    read_community = $7,
    write_community = $8,
    trap_community = $9,
    timeout = $10,
    poll_interval = $11,
    retries = $12
WHERE template_id = $5;

-- name: UpdateSnmpTemplateWithV3Parameters :exec
UPDATE snmp_template
SET name = $1,
    protocol = $2,
    version = $3,
    description = $4,
    updated_at = now()
WHERE id = $5;

UPDATE snmpv3_parameters
SET port = $6,
    security_level = $7,
    auth_protocol = $8,
    auth_password = $9,
    priv_protocol = $10,
    priv_password = $11,
    v3_user = $12,
    engine_id = $13,
    timeout = $14,
    poll_interval = $15,
    retries = $16
WHERE template_id = $5;

-- name: DeleteSnmpTemplateWithParameters :exec
DELETE FROM snmpv2_parameters
WHERE template_id = $1;

DELETE FROM snmpv3_parameters
WHERE template_id = $1;

DELETE FROM snmp_template
WHERE id = $1;

-- name: GetSnmpTemplateWithParameters :one
SELECT 
  t.*,
  v2.*,
  v3.*
FROM snmp_template t
LEFT JOIN snmpv2_parameters v2 ON v2.template_id = t.id
LEFT JOIN snmpv3_parameters v3 ON v3.template_id = t.id
WHERE t.id = $1;

-- name: ListSnmpTemplatesWithParameters :many
SELECT 
  t.*,
  v2.*,
  v3.*
FROM snmp_template t
LEFT JOIN snmpv2_parameters v2 ON v2.template_id = t.id
LEFT JOIN snmpv3_parameters v3 ON v3.template_id = t.id
ORDER BY t.id
LIMIT $1
OFFSET $2;