# Recover [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/dre1080/recover) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/dre1080/recover/master/LICENSE)

Recover is a HTTP middleware that catches any panics and serves a proper error response.

Works with all frameworks that support native http handler (eg. [Echo](https://github.com/labstack/echo), [Goji](https://github.com/zenazn/goji)).

## Installation

```
$ go get github.com/dre1080/recover
```

## Usage

``` go
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
```

Terminal

![Terminal](https://github.com/dre1080/recover/raw/master/img/screenshot-terminal.png)

HTML

![HTML](https://github.com/dre1080/recover/raw/master/img/screenshot-html.png)

JSON

![JSON](https://github.com/dre1080/recover/raw/master/img/screenshot-json.png)

### Echo Example

```go
...
func main() {
    e := echo.New()
    e.Use(recover.New(&recover.Options{
        Log: log.Print,
    }))

    ...

    e.Run(":3000")
}
```

### Goji Example

```go
...
func main() {
    goji.Use(recover.New(&recover.Options{
        Log: log.Print,
    }))

    ...

    goji.Serve()
}
```
