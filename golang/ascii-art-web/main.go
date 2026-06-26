package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// PageData holds the variables we pass to our HTML template
type PageData struct {
	InputText      string
	AsciiArt       string
	Error          string
	AvailableFonts []string
	selectedBanner   string
}

var (
	// A thread-safe cache for multiple fonts
	fontCache = make(map[string]map[rune][]string)
	cacheMu   sync.RWMutex
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleHome)

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	// 1. Get all available fonts from the 'fonts/' directory
	entries, err := os.ReadDir("fonts")
	if err != nil {
		data.Error = "Could not read fonts directory. Make sure it exists!"
	} else {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
				data.AvailableFonts = append(data.AvailableFonts, entry.Name())
			}
		}
	}

	// Handle Form Submission
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		selectedBanner := r.FormValue("font")
		
		data.InputText = text
		data.selectedBanner = selectedBanner

		if selectedBanner == "" {
			selectedBanner = "standard.txt" // Default fallback
		}

		// 2. Check if font is in cache (Read Lock)
		cacheMu.RLock()
		currentFont, ok := fontCache[selectedBanner]
		cacheMu.RUnlock()

		// 3. If not in cache, load it from disk and add to cache (Write Lock)
		if !ok {
			loadedFont, err := LoadBanner("banner/" + selectedBanner)
			if err != nil {
				data.Error = "Error loading font '" + selectedBanner + "': " + err.Error()
			} else {
				cacheMu.Lock()
				fontCache[selectedBanner] = loadedFont
				cacheMu.Unlock()
				currentFont = loadedFont
			}
		}

		// 4. Render the text if we successfully have a font
		if currentFont != nil {
			data.AsciiArt = RenderLine(text, currentFont)
		}
	}

	// Parse and execute the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}