package main

import (
	"html/template"
	"log"
	"net/http"
	pkgurl "net/url"

	"github.com/google/uuid"
)

type url struct {
	URL string
}

type urlhash map[string]string

var tmpl *template.Template

//var data string

func (u url) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var data string
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	if len(req.Form) != 0 && req != nil {
		data = GeneratetinyURL(req.Form)
	}
	tmpl.ExecuteTemplate(w, "tinyURL.gohtml", data)
}

func GeneratetinyURL(fv pkgurl.Values) string {
	TinyUrlKey := uuid.New().String()
	keytourlmap := urlhash{}
	keytourlmap[TinyUrlKey] = fv["url"][0]

	return TinyUrlKey
}

func init() {
	tmpl = template.Must(template.ParseFiles("tinyURL.gohtml"))
}

func main() {
	var u url
	http.ListenAndServe(":8080", u)

}
