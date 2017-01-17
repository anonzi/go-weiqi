
package config

import (
  "fmt"
  "net/http"
  "encoding/json"
  "go-weiqi/app/controllers"
)

type Route struct {
}

func InitRoute() *Route {
  route := &Route{}
  return route
}

type BaseJsonBean struct {
  Code    int         `json:"code"`
  Data    interface{} `json:"data"`
  Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
  return &BaseJsonBean{}
}

func (p *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" {
    controller.sayhelloName(w, r)
    return
  }
  http.NotFound(w, r)
  return
}

