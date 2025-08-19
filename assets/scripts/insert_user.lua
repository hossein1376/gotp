-- KEYS[1] = users prefix key + phone
-- ARGV[1] = phone number
-- ARGV[2] = created at
-- ARGV[3] = last login

local key = KEYS[1]
local phone = ARGV[1]
local created = tostring(ARGV[2])
local last_login = tostring(ARGV[2])

if redis.call("EXISTS", key) == 0 then
    redis.call(
            "HSET", key,
            "phone", phone,
            "created_at", created,
            "last_login", last_login
    )
    return
else
    redis.call("HSET", key, "last_login", last_login)
end