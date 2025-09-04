package main

import (
	"gnomik/internal/api"
	"log"
	"net/http"
)

func main() {
	// 1. Создаем обработчики
	handler := api.NewHandler()

	// 2. Регистрируем роуты (ДО запуска сервера!)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/convert", handler.Convert)

	// 3. Статические файлы (ДО запуска сервера!)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 4. Логируем ПЕРЕД запуском
	log.Println("Server started on port 8080")

	// 5. Запускаем сервер (ОДИН РАЗ!)
	log.Fatal(http.ListenAndServe(":8080", nil))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Относительный путь к файлу
	http.ServeFile(w, r, "static/index.html")
}
