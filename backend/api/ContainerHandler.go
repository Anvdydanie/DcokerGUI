package api

import (
	"DockerGUI/backend/api/containerActions"
	"errors"
	"github.com/docker/docker/client"
	"net/http"
)

func ContainerHandler(w http.ResponseWriter, req *http.Request) {
	var result []byte
	var path = req.Method + " " + req.URL.Path
	var params = req.URL.Query()
	//var body = req.Body

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	switch path {
	// метод выводит список контейнеров
	// http://localhost:9999/api/containers/list?showAll= 1 или 0
	case "GET /api/containers/list":
		err, result = containerActions.List(cli, params)
	// метод создает новый контейнер
	case "POST /api/containers/create":
		err, result = containerActions.Create(cli, params) // TODO недоделан
	// получение информации о контейнере
	// http://localhost:9999/api/containers/list?id= id_контейнера
	case "GET /api/containers/inspect":
		err, result = containerActions.Inspect(cli, params)
	// получение списка процессов, протекающих в контейнере
	case "GET /api/containers/processes":
		err, result = containerActions.Processes(cli, params)
	// получение логов контейнера
	case "GET /api/containers/logs":
		err, result = containerActions.Logs(cli, params)
	// получение информации по использованию ресурсов контейнера
	case "GET /api/containers/stats":
		err, result = containerActions.Stats(cli, params) // TODO протестировтаь
	// метод запуска контейнера
	case "POST /api/containers/start":
		err, result = containerActions.Start(cli, params)
	// метод запуска контейнера
	case "POST /api/containers/start-all":
		err, result = containerActions.StartAll(cli)
	// метод останавливает контейнер
	case "POST /api/containers/stop":
		err, result = containerActions.Stop(cli, params)
	// останавливает все контейнеры
	case "POST /api/containers/stop-all":
		err, result = containerActions.StopAll(cli)
	// перезапуск контейнера
	case "POST /api/containers/restart":
		err, result = containerActions.Restart(cli, params)
	// принудительно убиваем контейнер
	case "POST /api/containers/kill":
		err, result = containerActions.Kill(cli, params)
	// удаление контейнера
	case "DELETE /api/containers/remove":
		err, result = containerActions.Remove(cli, params)
	default:
		err, result = errors.New("метод "+path+" не найден в API"), nil
	}

	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(result)
}
