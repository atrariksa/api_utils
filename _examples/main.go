package main

import (
	"log"
	"net/http"

	"github.com/atrariksa/api_utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ExampleXMLHandler struct {
	api_utils.DefaultHttpHandler
}

func (xh *ExampleXMLHandler) Handle(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `xml:"message"`
	}
	resp := xh.Process(r.Context(), nil)
	xh.Write(w, 200, response{Message: resp.(string)})
}

func setupApis() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	dh := api_utils.GetDefaultHandler()
	xmlh := ExampleXMLHandler{}
	xmlh.DefaultHttpHandler = dh
	xmlh.IRespWriter = &api_utils.XmlRespWriter{}

	r.Get("/", dh.Handle)
	r.Get("/xml", xmlh.Handle)
	return r
}

func main() {
	server := &http.Server{Addr: "0.0.0.0:4567", Handler: setupApis()}
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
