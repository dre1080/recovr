/*Package recover is a HTTP middleware that catches any panics and serves a proper error response.

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

A GET request to "/" will output:

        [PANIC] you should not have a handler that just panics ;)

        Stack Trace:
        ============

        path/to/main.go:11                                              glob.func1
        /usr/local/Cellar/go/1.5.1/libexec/src/net/http/server.go:1423  HandlerFunc.ServeHTTP
        github.com/dre1080/recover/recover.go:66                        New.func1.1
        /usr/local/Cellar/go/1.5.1/libexec/src/net/http/server.go:1423  HandlerFunc.ServeHTTP
        /usr/local/Cellar/go/1.5.1/libexec/src/net/http/server.go:1863  serverHandler.ServeHTTP
        /usr/local/Cellar/go/1.5.1/libexec/src/net/http/server.go:1362  (*conn).serve
        /usr/local/Cellar/go/1.5.1/libexec/src/runtime/asm_amd64.s:1697 goexit
*/
package recover
