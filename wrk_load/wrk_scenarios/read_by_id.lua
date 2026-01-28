local POOL_START = 1
local POOL_END = 100000

math.randomseed(os.time())

request = function()
	local id = POOL_START + math.random(0, POOL_END - POOL_START)
	local path = "/objects/" .. id

	return wrk.format("GET", path, {}, nil)
end
