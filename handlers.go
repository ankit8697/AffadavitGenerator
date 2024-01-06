package main

import (
	"github.com/nguyenthenguyen/docx"
	"net/http"
	"log"
	"html/template"
)

func buildS6Affadavit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Fatal("server recieved non-POST request")
		return
	}
	name := r.FormValue("name")
	parentName := r.FormValue("parentName")
	age := r.FormValue("age")
	address := r.FormValue("address")
	doc, err := docx.ReadDocxFile("./Affidavit under s.6.docx")
	defer doc.Close()
	
	if err != nil {
		log.Fatal("the Affadavit file for section 6 could not be read")
	}
	affadavit := doc.Editable()
	affadavit.Replace("FIRST_NAME", name, -1)
	affadavit.Replace("PARENT_NAME", parentName, -1)
	affadavit.Replace("USER_AGE", age, -1)
	affadavit.Replace("USER_ADDRESS", address, -1)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	affadavit.Write(w)
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	// Create an instance of the template with your HTML
	tmpl := template.Must(template.ParseFiles("index.html"))

	// Execute the template, passing in the data
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}