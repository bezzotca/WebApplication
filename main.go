package main

import (
	"html/template"
)

// Парсится один раз при старте: если файла нет, программа упадёт сразу
// с понятным сообщением, а не паникой на первом запросе.
var mainPage = template.Must(template.ParseFiles("Pages/HTML/Bootstrap/MainPage.html"))

func main() {
	handleRequest()
}
