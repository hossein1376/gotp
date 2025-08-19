-- KEYS[1] = bucket name
-- ARGV[1] = credit
-- ARGV[2] = cost
-- ARGV[3] = window

local bucketName = KEYS[1]
local credit = tonumber(ARGV[1])
local cost = tonumber(ARGV[2])
local window = tonumber(ARGV[3])

local tokens = tonumber(redis.call("GET", bucketName))

if tokens == nil then
    tokens = credit
    redis.call("SET", bucketName, tokens, "PX", window)
elseif tokens == 0 then
    return 0
end

local isAllowed = 0
if tokens >= cost then
    isAllowed = 1
    tokens = tokens - cost
else
    tokens = 0
end
redis.call("SET", bucketName, tokens, "KEEPTTL")

return isAllowed
