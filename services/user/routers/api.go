package routers

import (
	"net/http"
	"user/controllers"
)

func Home() {
	url := "/api/v1/user"
	http.HandleFunc(url+"/login", controllers.Login)
}
