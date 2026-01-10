math.randomseed(os.time())

local function random_float(lower, upper)
	return lower + math.random() * (upper - lower)
end

request = function()
	local id = math.random(1, 1000000)
	local lat = random_float(-90, 90)
	local lon = random_float(-180, 180)
	local time = os.date("!%Y-%m-%dT%H:%M:%SZ")

	local body = string.format('{"id": %d, "latitude": %.6f, "longitude": %.6f, "time": "%s"}', id, lat, lon, time)

	local method = "POST"
	local path = "/objects"
	local headers = {}
	headers["Content-Type"] = "application/json"

	return wrk.format(method, path, headers, body)
end
