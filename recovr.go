package recovr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ztrue/tracerr"
)

const (
	applicationJSON            = "application/json"
	applicationJSONCharsetUTF8 = applicationJSON + "; charset=utf-8"
)

var contentType = http.CanonicalHeaderKey("Content-Type")

type jsonResponse struct {
	Message  string `json:",omitempty"`
	Location string `json:",omitempty"`
}

// New returns a middleware that: recovers from panics, logs the panic and backtrace,
// writes a HTML response and Internal Server Error status.
//
// If a JSON content type is detected it returns a JSON response.
func New() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					err, ok := rec.(error)
					if !ok {
						err = fmt.Errorf("%v", rec)
					}

					e := tracerr.Wrap(err)
					frames := e.StackTrace()[2:]

					tracerr.PrintSourceColor(tracerr.CustomError(err, frames))

					w.WriteHeader(http.StatusInternalServerError)

					ct := r.Header.Get(contentType)
					if strings.HasPrefix(ct, applicationJSON) {
						w.Header().Set(
							contentType,
							applicationJSONCharsetUTF8,
						)

						e := jsonResponse{
							Message:  err.Error(),
							Location: frames[0].String(),
						}
						json.NewEncoder(w).Encode(e)
					} else {
						w.Write(compileTemplate(r, err, frames))
					}
				}
			}()

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
