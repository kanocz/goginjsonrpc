package main

import (
  "github.com/kanocz/goginjsonrpc"
  "github.com/gin-gonic/gin"
)

type TestRPC struct {
  counter int
}

func (t *TestRPC) Add(arg int) int {
  t.counter += arg
  return t.counter
}

func (t *TestRPC) Sub(arg int) int {
  t.counter += arg
  return t.counter
}

func main() {

  rpc := TestRPC{}

  router := gin.Default()
  router.POST("/", func(c *gin.Context) { goginjsonrpc.ProcessJsonRPC(c, &rpc); })

  router.Run("127.0.0.1:8000")
}
