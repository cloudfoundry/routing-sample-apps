# http2-sample

## Server:

Run the Server locally:

`PORT=8765 go run main.go`

You can also push as a cf app.

## Client:

You should run the client locally.

If running against a local server:
```
go run main.go -host="127.0.0.1:8576"
```

If running against a remote env:
```
go run main.go -host="http2-server.dev-full-2.routing.cf-app.com" -routerAddr="envoy.istio.dev-full-2.routing.cf-app.com"
```

Verify envoy stdout logs:
```
[2018-06-01T22:46:22.192Z] "GET / HTTP/2" 200 - 0 6 1 1 "204.244.11.130" "Go-http-client/2.0" "816fbc1e-fc69-9f11-b845-5b3dda9c0391" "http2-server.dev-full-2.routing.cf-app.com" "10.0.32.15:61000"
```

