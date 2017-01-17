

package main

import (
  "net/http"
  "go-weiqi/config"
)

func main() {
  route := config.InitRoute()
  http.ListenAndServe(":9090", route)
}
