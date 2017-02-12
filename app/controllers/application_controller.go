
package controller

import (
  "fmt"
  "net/http"
  // "go-weiqi/global"
  // "github.com/adjust/gorails/marshal"
)

type ApplicationController struct {
  w http.ResponseWriter
  r *http.Request
}

func (controller *ApplicationController)CurrentUserId() int64{
  var currentId int64
  fmt.Println(controller.r.Cookie)
  // cookie, _ := controller.r.Cookie("_Weiqi_session")
  // if cookie != nil {
  //   val, _ := global.SessionRedis.Get("weiqi:session:"+cookie.Value)
  //     session_data, err := marshal.CreateMarshalledObject(val).GetAsMap()
  //     if err == nil {
  //       uid, err := session_data["user_id"].GetAsInteger()
  //         if err == nil {
  //           fmt.Println(uid)
  //           currentId = uid
  //         }
  //     }
  // }
  return currentId
}
