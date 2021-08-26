# go-rpc-example

## Dependencies
	📍 github.com/gorilla/mux 
	📍 github.com/gorilla/rpc 
## 🚀 Run Server
```go run jsonRpcServer/main.go```

## 🚀 GET Request 

```
    curl -X POST \
    http://localhost:1234/rpc \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{
    "method" : "JSONServer.GiveBookDetail",
    "params" : [{
    "ID": "3"
    }],
    "id": "1"
}'
```

