
package controller

type HomeController struct {
  ApplicationController
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form["aa"])
  result := NewBaseJsonBean()
  result.Code = 100
  result.Message = "登录成功"
  bytes, _ := json.Marshal(result)
  fmt.Fprint(w, string(bytes))
}

