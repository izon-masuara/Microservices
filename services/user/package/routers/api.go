package routers

import (
	"user/package/controllers"

	"github.com/gorilla/mux"
)

// All of routers
func ApiUser(http *mux.Router) {
	url := "/api/v1/user"
	http.HandleFunc(url+"/login", controllers.Login)
	http.HandleFunc(url+"/register", controllers.Register)
	http.HandleFunc(url+"/token", controllers.CheckToken)
}
