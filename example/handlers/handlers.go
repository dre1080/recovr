package handlers

import "net/http"

func PanicHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("you should not have a handler that just panics ;)")
	})
}
