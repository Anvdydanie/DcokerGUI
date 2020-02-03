package main

import (
	"DockerGUI/backend/api"
	"fmt"
	"net/http"
)

func main() {
	// Обработка запросов по контейнерам
	http.HandleFunc("/api/containers/", api.ContainerHandler)
	// Обработка запросов по образам
	http.HandleFunc("/api/images/", api.ImageHandler)

	// запускаем вебсервер
	err := http.ListenAndServe(":9999", nil)
	if err == nil {
		fmt.Println("Server is listening")
	} else {
		fmt.Println(err)
	}
}
