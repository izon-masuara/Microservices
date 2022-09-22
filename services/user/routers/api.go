package routers

import "net/http"

func Home() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Home"))
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("register"))
	})
}
