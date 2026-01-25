package app

import "net/http"

func Start() {
	http.ListenAndServe(":8080", nil)
}
