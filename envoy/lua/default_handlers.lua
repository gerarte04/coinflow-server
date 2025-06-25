socket = require "socket"

start, authority, path, method = nil, nil, nil, nil

function envoy_on_request(request_handle)
    local headers = request_handle:headers()
    authority = headers:get(":authority")
    path = headers:get(":path")
    method = headers:get(":method")

    start = socket.gettime()
end

function envoy_on_response(response_handle)
    local headers = response_handle:headers()
    local code_header = headers:get("x-http-code")

    if code_header ~= nil then
        headers:replace(":status", code_header)
    end

    response_handle:logInfo(string.format("[HTTP/1.1] %s | %s | %s | %s | %.5fs",
        authority, path, method, headers:get(":status"), socket.gettime() - start
    ))
end
