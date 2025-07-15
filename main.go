package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name            string
	Description     string
	DurationInHours string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", CourseHandler)

	http.ListenAndServe(":8090", mux)
}

func CourseHandler(responseWriter http.ResponseWriter, request *http.Request) {
	courses := []Course{
		{
			Name:            "Curso 1",
			Description:     "Descrição do Curso 1",
			DurationInHours: "40",
		},
		{
			Name:            "Curso 2",
			Description:     "Descrição do Curso 2",
			DurationInHours: "30",
		},
	}

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	courseTemplate := template.Must(template.New("content.html").ParseFiles(templates...))
	err := courseTemplate.ExecuteTemplate(responseWriter, "content.html", courses)
	if err != nil {
		http.Error(responseWriter, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
