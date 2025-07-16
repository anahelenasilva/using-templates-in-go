package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Course struct {
	Name            string
	Description     string
	DurationInHours string
}

func FormatDuration(duration string) string {
	if duration == "" {
		return "N/A"
	}
	return fmt.Sprintf("%s horas", duration)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", CourseHandler)

	http.ListenAndServe(":8090", mux)
}

// NotFoundHandler handles all non-existent routes
func NotFoundHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	responseWriter.Header().Set("Content-Type", "text/html")

	// Create data to pass to the template
	data := struct {
		URL string
	}{
		URL: request.URL.Path,
	}

	// Parse and execute the 404 template
	notFoundTemplate := template.Must(template.ParseFiles("404.html"))
	err := notFoundTemplate.Execute(responseWriter, data)
	if err != nil {
		// Fallback to simple text if template fails
		http.Error(responseWriter, "404 - Page Not Found", http.StatusNotFound)
		return
	}
}

func CourseHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// Only handle exact root path, redirect others to 404
	if request.URL.Path != "/" {
		NotFoundHandler(responseWriter, request)
		return
	}

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

	courseTemplate := template.New("content.html")
	courseTemplate.Funcs(template.FuncMap{
		"FormatDuration": FormatDuration,
	})
	courseTemplate = template.Must(courseTemplate.ParseFiles(templates...))

	err := courseTemplate.ExecuteTemplate(responseWriter, "content.html", courses)
	if err != nil {
		http.Error(responseWriter, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
