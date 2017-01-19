
package controller

import (
  "bytes"
  "fmt"
  "net/http"
  "encoding/gob"
  "encoding/json"
  "html/template"
  "go-weiqi/global"
  // "strings"
  // "io"
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

func NewHomeController() *HomeController {
  return &HomeController{}
}

func (controller *HomeController)Index(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form["aa"])
  controller.CurrentUser()
  result := NewBaseJsonBean()
  result.Code = 100
  result.Message = "登录成功"
  bytes, _ := json.Marshal(result)
  fmt.Fprint(w, string(bytes))
}

func (controller *HomeController)Test(w http.ResponseWriter, r *http.Request) {
  cookie, _ := r.Cookie("_seat_session")
  val, _ := global.SessionRedis.Get("seat:session:"+cookie.Value)
  fmt.Println(string(val))

  m := make(map[string]interface{})
	err := json.Unmarshal(val, &m)
	if err != nil {
		fmt.Printf("redistore.JSONSerializer.deserialize() Error: %v", err)
	}
  fmt.Println(m)


  dec := gob.NewDecoder(bytes.NewBuffer(val))
  dec.Decode(&m)
  fmt.Println(m)


  t, _ := template.ParseFiles("app/views/home/test.gtpl")
  a := "aa"
  t.Execute(w, a)
}

