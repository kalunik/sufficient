package controller

import (
	model "app/models"
	"fmt"
	"net/http"
)

func prepareData(uid string, cache model.Cache, w http.ResponseWriter) string {
	data, ok := cache.Get(uid)
	if !ok {
		return "There is no data with associated order_uid " + uid
	}
	return string(data.([]byte))
}

func HttpHandler(cache model.Cache) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.Error(w, "404 not found. ", http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./static/form.html")
		case "POST":

			fmt.Fprintf(w, "<a href=\"/\">Home page</a>")

			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "<div>ParseForm() err: %v</div>", err)
				return
			}
			uid := r.FormValue("order_uid")
			fmt.Fprintln(w, "<div>", prepareData(uid, cache, w), "</div>")

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}
