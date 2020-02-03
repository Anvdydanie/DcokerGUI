package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод выводит статистику затрат ресурсов контейнера */
func Stats(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	containerStats, err := cli.ContainerStats(context.Background(), id, false)
	if err == nil {
		result, err = json.Marshal(containerStats)
	}

	return err, result
}
