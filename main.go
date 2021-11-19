package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//go:embed assets/*
var assets embed.FS

//go:embed templates
var templates embed.FS

type data struct {
	Spec string
}

func main() {
	spec := flag.String("spec", "./spec.yml", "a url or path to the openapi spec")
	doc := flag.String("doc", "swagger", "documentation type (swagger or redoc)")
	port := flag.Int("port", 9000, "port to serve http over")
	host := flag.String("host", "0.0.0.0", "host ip to serve using")

	flag.Parse()

	templates := map[string]*template.Template{
		"swagger": template.Must(template.ParseFS(templates, "templates/swagger.html")),
		"redoc":   template.Must(template.ParseFS(templates, "templates/redoc.html")),
	}

	tmpl := templates["swagger"]
	if v, ok := templates[*doc]; ok {
		tmpl = v
	}

	mux := http.NewServeMux()

	data, bytes, err := prepareData(*spec)
	if err != nil {
		log.Fatal(err)
	}

	// Spec is loaded from a file path
	if len(bytes) != 0 {
		mux.HandleFunc("/"+data.Spec, func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write(bytes); err != nil {
				log.Fatal(err)
			}
		})
	}

	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, data); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Listening on http://%s:%d", *host, *port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), mux); err != nil {
		log.Fatal(err)
	}
}

func prepareData(spec string) (data, []byte, error) {
	if _, err := url.ParseRequestURI(spec); err == nil {
		return data{
			Spec: spec,
		}, []byte{}, nil
	}

	if _, err := os.Stat(spec); os.IsNotExist(err) {
		return data{}, []byte{}, fmt.Errorf("swagger file not found for path %s", spec)
	}

	bytes, err := ioutil.ReadFile(spec)
	if err != nil {
		return data{}, []byte{}, err
	}

	d := data{
		Spec: filepath.Base(spec),
	}

	return d, bytes, nil
}
