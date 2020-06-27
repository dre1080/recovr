/*Package recovr is a HTTP middleware that catches any panics and serves a proper error response.

package main

import (
	"log"
	"net/http"

	"github.com/dre1080/recovr"
)

var myPanicHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	panic("you should not have a handler that just panics ;)")
})

func main() {
	recovery := recovr.New()
	app := recovery(myPanicHandler)
	http.ListenAndServe("0.0.0.0:3000", app)
}
*/
package recovr
