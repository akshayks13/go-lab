package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files (HTML)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Handle form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			age := r.FormValue("age")
			fmt.Fprintf(w, "Received: Name = %s, Age = %s", name, age)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
