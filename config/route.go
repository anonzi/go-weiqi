
package config

import (
  "net/http"
  "go-weiqi/app/controllers"
)

type Route struct {
}

func NewRoute() *Route {
  route := &Route{}
  return route
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  homeController := controller.NewHomeController()
  if r.URL.Path == "/" {
    homeController.Index(w, r)
    return
  }
  if r.URL.Path == "/test" {
    homeController.Test(w, r)
    return
  }
  http.NotFound(w, r)
  return
}

