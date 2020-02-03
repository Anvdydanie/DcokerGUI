package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод запускает контейнер */
func Start(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	if err = cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err == nil {
		result, err = json.Marshal("Контейнер " + id + " запущен")
	}

	return err, result
}
