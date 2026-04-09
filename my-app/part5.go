package main

import (
	"fmt"
	//"net/http"
	"html/template"
	"log"
	"log/slog"
	"os"
)

func Part5() {
	tmpl := `{{.}}`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	t2 := template.Must(template.New("").ParseGlob("template/*.tmpl"))
	err2 := t2.ExecuteTemplate(os.Stdout, "index", "これは本文です")
	if err2 != nil {
		log.Fatal(err2)
	}

	slog.Info("test", "Level", "Info")
	slog.Debug("test", "Level", "Debug")
	slog.Warn("test", "Level", "Warn")
	slog.Error("test", "Level", "Error")

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello World")
	//})
	//http.ListenAndServe(":8080", nil)
}
