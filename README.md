# go-gin-jsonrpc
Very basic and simple jsonrpc server implementation for golang and gin-gonic

## Example of usage:

```go
  rpc := TestRPC{}
  router := gin.Default()
  router.POST("/", func(c *gin.Context) { goginjsonrpc.ProcessJsonRPC(c, &rpc); })
  router.Run("127.0.0.1:8000")
```
