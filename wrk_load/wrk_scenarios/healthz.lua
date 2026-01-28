request = function()
	return wrk.format("GET", "/healthz", {}, nil)
end
