local POOL_START = 1
local POOL_END = 100000

math.randomseed(os.time())

request = function()
	local id = POOL_START + math.random(0, POOL_END - POOL_START)
	local lat = -90 + math.random() * 180
	local lon = -180 + math.random() * 360
	local status = math.random() > 0.5
	local time = os.date("!%Y-%m-%dT%H:%M:%SZ")

	local body = string.format(
		'{"id": %d, "status": %s, "latitude": %.6f, "longitude": %.6f, "time": "%s"}',
		id,
		tostring(status),
		lat,
		lon,
		time
	)

	local headers = {}
	headers["Content-Type"] = "application/json"

	return wrk.format("POST", "/objects", headers, body)
end
