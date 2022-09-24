package routers

import (
	"net/http"
	"user/controllers"
)

func ApiUser() {
	url := "/api/v1/user"
	http.HandleFunc(url+"/login", controllers.Login)
}
