function envoy_on_request(request_handle)
end

function envoy_on_response(response_handle)
    local headers = response_handle:headers()
    local code_header = headers:get("x-http-code")

    if code_header ~= nil then
        headers:replace(":status", code_header)
    end
end
