package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()

	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	pr.Add("/", homePage)

	http.ListenAndServe(":8080", pr) // Pass pr as the handler
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{
		make(map[string]http.HandlerFunc),
	}
}

func (p *pathResolver) Add(pattern string, handler http.HandlerFunc) {
	p.handlers[pattern] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
			return
		}
	}

	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "`default value:` Inigo Montoya"
	}
	fmt.Fprintf(res, "Hello, my name is %s", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	parts := strings.Split(path, "/")

	name := parts[2]
	if name == "" {
		name = "`default value:` Inigo Montoya"
	}
	fmt.Fprintf(res, "Goodbye %s", name)
}
