
package controller

import (
  "fmt"
  "net/http"
  "encoding/json"
  "html/template"
)

type HomeController struct {
  ApplicationController
}

type BaseJsonBean struct {
  Code    int         `json:"code"`
  Data    interface{} `json:"data"`
  Message string      `json:"message"`
}


func NewBaseJsonBean() *BaseJsonBean {
  return &BaseJsonBean{}
}

func NewHomeController(w http.ResponseWriter, r *http.Request) *HomeController {
  return &HomeController{ApplicationController{w: w, r: r}}
}

func (controller *HomeController)Index() {
  controller.r.ParseForm()
  fmt.Println(controller.r.Form["aa"])
  result := NewBaseJsonBean()
  result.Code = 100
  result.Message = "登录成功"
  bytes, _ := json.Marshal(result)
  fmt.Fprint(controller.w, string(bytes))
}

func (controller *HomeController)Test() {
  t, _ := template.ParseFiles("app/views/home/test.gtpl")
  a := "aa"
  fmt.Println(controller.CurrentUserId())
  t.Execute(controller.w, a)
}

