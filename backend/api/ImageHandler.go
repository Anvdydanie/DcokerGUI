package api

import (
	"DockerGUI/backend/api/imageActions"
	"github.com/docker/docker/client"
	"net/http"
)

func ImageHandler(w http.ResponseWriter, req *http.Request) {
	var result []byte
	var path = req.Method + " " + req.URL.Path
	var params = req.URL.Query()
	//var body = req.Body

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	switch path {
	// метод выводит список названий образов
	// http://localhost:9999/api/images/list?showAll= 1 или 0
	case "GET /api/images/list":
		err, result = imageActions.List(cli, params)
	// метод создает новый образ из dokerfile
	case "POST /api/images/build":
		err, result = imageActions.Build(cli, params) // TODO доделать
	// Create an image either by pulling it from the registry or by importing it
	case "POST /api/images/create":
		err, result = imageActions.Create(cli, params) // TODO проверить
	// Pull
	case "POST /api/images/pull":
		err, result = imageActions.Pull(cli, params)
	// Push
	case "POST /api/images/push":
		err, result = imageActions.Push(cli, params) // TODO доделать
	// метод удаляет образ
	case "DELETE /api/images/remove":
		err, result = imageActions.Remove(cli, params) // TODO проверить
	}

	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(result)
}
