package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", UpdateData)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var address = "localhost:8080"
	log.Println("server started at ", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func UpdateData(w http.ResponseWriter, r *http.Request) {

	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	water := rand.Intn(100)
	var status_water string
	var status_wind string
	if water < 5 {
		status_water = "aman"
	} else if water > 5 && water <= 8 {
		status_water = "siaga"
	} else {
		status_water = "bahaya"
	}

	wind := rand.Intn(100)
	if wind < 6 {
		status_wind = "aman"
	} else if wind > 6 && water <= 15 {
		status_wind = "siaga"
	} else {
		status_wind = "bahaya"
	}

	var message1 = fmt.Sprintf("water %d\n", water)
	var message2 = fmt.Sprintf("wind %d", wind)

	var data = map[string]interface{}{
		"water":        message1,
		"wind":         message2,
		"status_water": status_water,
		"status_wind":  status_wind,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
