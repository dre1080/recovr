package main

import (
	"log"
	"net/http"

	"github.com/dre1080/recover"
)

var myPanicHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	panic("you should not have a handler that just panics ;)")
})

func main() {
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})

	// recoveryWithDefaults := recovery.New(nil)

	app := recovery(myPanicHandler)
	http.ListenAndServe("0.0.0.0:3000", app)
}
