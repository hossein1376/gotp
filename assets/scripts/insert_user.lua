-- key[1] users prefix key + phone
-- arg[1] phone number
-- arg[2] created at
-- arg[3] last login

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