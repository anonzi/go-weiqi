
package config

import (
  "net/http"
  "go-weiqi/app/controllers"
  "fmt"
)

type Route struct {
}

func NewRoute() *Route {
  route := &Route{}
  return route
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  homeController := controller.NewHomeController(w, r)
  fmt.Println("Start Request")
  defer fmt.Println("End Request")
  if r.URL.Path == "/" {
    homeController.Index()
    return
  }
  if r.URL.Path == "/test" {
    homeController.Test()
    return
  }
  http.NotFound(w, r)
  return
}

