math.randomseed(os.time())

request = function()
	local random_id = math.random(1, 1000)

	local path = "/" .. random_id

	return wrk.format("GET", path)
end
