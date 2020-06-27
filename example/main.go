package main

import (
	"net/http"

	"github.com/dre1080/recovr"
	"github.com/dre1080/recovr/example/handlers"
)

func main() {
	recovery := recovr.New()
	app := recovery(handlers.PanicHandler())
	http.ListenAndServe("0.0.0.0:3000", app)
}
