package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if err := mainPage.Execute(w, nil); err != nil {
		log.Printf("render MainPage: %v", err)
	}
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bbb")
}

func disciplinesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ccc")
}

func handleRequest() {
	// Раздача файлов из папки Images: /Images/Logos/x.png -> Images/Logos/x.png
	http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("Images"))))

	// "/{$}" — только корень. Без {$} шаблон "/" перехватывает все пути,
	// и отсутствующие файлы возвращают HTML со статусом 200 вместо 404.
	http.HandleFunc("/{$}", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.HandleFunc("/disciplines/", disciplinesPage)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
