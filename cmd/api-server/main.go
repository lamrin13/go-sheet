package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-sheet/internal"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {

	id := "1nU6SClw4Q66FWDaaJrAuf1gvFDQdm85oXHwZCDDRxho"
	columns := "!A:E"

	flag.Parse()

	r := chi.NewRouter()

	corsConfig := cors.New(cors.Options{
		// AllowedOrigins is a list of origins a cross-domain request can be executed from
		// You can allow all origins with "*"
		AllowedOrigins: []string{"https://*", "http://*"}, // Allow only this origin
		// AllowedMethods is a list of methods the client is allowed to use
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders is a list of non-simple headers the client is allowed to use
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		// AllowCredentials should be set to true if you allow cookies
		AllowCredentials: true,
	})

	r.Use(corsConfig.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
		body, err := internal.GetAll(id)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.Write(body)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

		var rows []internal.Request
		fmt.Println(r.Body)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&rows)
		if err != nil {
			fmt.Println("Error while decoding json ", err)
		}
		body, err := internal.Append(columns, id, rows)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.Write(body)
	})

	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		body, err := internal.GetCategories(id)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Write(body)
	})

	http.ListenAndServe(":8080", r)
	print("server started")
}
