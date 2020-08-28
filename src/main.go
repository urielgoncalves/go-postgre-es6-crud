package main

import (
	Entities "domain/entities"
	"html/template"
	"net/http"
)

type CustomerPageData struct {
	PageTitle string
	Customers []Entities.Customer
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := CustomerPageData{
			PageTitle: "My Customers list",
			Customers: []Entities.Customer{
				{Firstname: "Uriel", Lastname: "Goncalves", Birthdate: "1985-04-09", Gender: "Male", Email: "urielgoncalves@teste.com", Address: "Abbey Street Upper, Dublin"},
				{Firstname: "John", Lastname: "Moore", Birthdate: "1995-04-01", Gender: "Male", Email: "john@teste.com", Address: "Abbey Street Upper, Dublin"},
				{Firstname: "Saoirse", Lastname: "O'Brien", Birthdate: "1999-09-09", Gender: "Female", Email: "saoirse@teste.com", Address: "Abbey Street Upper, Dublin"},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8000", nil)
}
