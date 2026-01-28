request = function()
	return wrk.format("GET", "/objects?status=true", {}, nil)
end
