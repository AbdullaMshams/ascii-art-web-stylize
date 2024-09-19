package main

import (
	"asciiartweb/arts"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var tmpl *template.Template
var result string

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/AsciiArt", handleSubmit)
	http.HandleFunc("/toggle-theme", handleToggleTheme)  // New handler for theme toggle

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}

// Serve the index page with theme support
func serveIndex(w http.ResponseWriter, r *http.Request) {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html", "templates/error404.html", "templates/error500.html", "templates/error400.html", "templates/error405.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parseErr := template.Must(template.ParseFiles("templates/error500.html"))
		parseErr.Execute(w, nil)
		return
	}

	// Handle 404 Not Found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.ExecuteTemplate(w, "error404.html", nil); err != nil {
			http.Error(w, "Error rendering 404 page", http.StatusNotFound)
		}
		return
	}

	theme := getThemeFromCookie(r)

	data := struct {
		Result string
		Theme  string
	}{
		Result: result,
		Theme:  theme,
	}

	// Serve the index template with theme support
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

// Handle form submission for ASCII art generation
func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		inputText := r.FormValue("inputText")
		input := strings.Split(inputText, "\r\n")
		style := r.FormValue("style")
		if style != "standard" && style != "shadow" && style != "thinkertoy" {
			if err := tmpl.ExecuteTemplate(w, "error400.html", nil); err != nil {
				http.Error(w, "Error rendering 400 page", http.StatusBadRequest)
			}
			return
		}
		for _, str := range input {
			for _, char := range str {
				if char < 0 || char > 127 {
					if err := tmpl.ExecuteTemplate(w, "error400.html", nil); err != nil {
						http.Error(w, "Error rendering 400 page", http.StatusBadRequest)
					}
					return
				}
			}
		}
		result = arts.Generator(input, style)
		http.Redirect(w, r, "/", http.StatusSeeOther)  // Redirect to reload with generated art
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := tmpl.ExecuteTemplate(w, "error405.html", nil); err != nil {
			http.Error(w, "Error rendering 405 page", http.StatusMethodNotAllowed)
		}
	}
}

// Handle theme toggle and set theme cookie
func handleToggleTheme(w http.ResponseWriter, r *http.Request) {
	theme := r.FormValue("theme")

	// Set the theme cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "theme",
		Value: theme,
		Path:  "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Get theme from cookie or default to light
func getThemeFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("theme")
	if err != nil {
		return "light" // Default theme
	}
	return cookie.Value
}
