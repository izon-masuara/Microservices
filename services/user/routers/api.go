package routers

import (
	"net/http"
	"user/controllers"
)

func ApiUser() {
	url := "/api/v1/user"
	http.HandleFunc(url+"/login", controllers.Login)
	http.HandleFunc(url+"/register", controllers.Register)
	http.HandleFunc(url+"/token", controllers.CheckToken)
}
