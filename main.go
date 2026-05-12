package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/projects.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/projects", projectsPage)
	http.HandleFunc("/contact", contactPage)

	log.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}