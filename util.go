package recovr

//go:generate go run github.com/gobuffalo/packr/v2/packr2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/ztrue/tracerr"
)

func mustReadFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func mustReadLines(filename string) []string {
	bytes := mustReadFile(filename)
	return strings.Split(string(bytes), "\n")
}

func getLines(frames []tracerr.Frame) []string {
	var lines []string
	for _, frame := range frames {
		src := mustReadLines(frame.Path)
		start := frame.Line - 6
		end := frame.Line + 5

		if start < 0 {
			start = 0
		}

		totalLines := len(src) - 1
		if end > totalLines {
			end = totalLines
		}

		lines = append(lines, strings.Join(src[start:end], "\n"))
	}
	return lines
}

func compileTemplate(r *http.Request, err error, frames []tracerr.Frame) []byte {
	if len(frames) == 0 {
		return nil
	}

	var buf bytes.Buffer

	box := packr.New("recovr", "./templates")
	html, boxErr := box.FindString("panic.html")
	if boxErr != nil {
		panic(boxErr)
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"sub": func(i int, sub int) int {
			return i - sub
		},
	}

	t := template.Must(template.New("recovr").Funcs(funcMap).Parse(html))
	t.Execute(&buf, struct {
		URL    string
		Err    error
		Lines  []string
		Frames []tracerr.Frame
	}{
		r.URL.Path,
		err,
		getLines(frames),
		frames,
	})

	return buf.Bytes()
}
