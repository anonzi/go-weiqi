
package main

import (
  "net/http"
  "go-weiqi/config"
  "go-weiqi/global"
)

func main() {
  route := config.NewRoute()
  global.SessionRedis.Db = 2
  http.ListenAndServe(":9090", route)
}
