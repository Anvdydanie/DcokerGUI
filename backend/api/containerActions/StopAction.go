package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод останавливает контейнер */
func Stop(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	if err = cli.ContainerStop(context.Background(), id, nil); err == nil {
		result, err = json.Marshal("Контейнер " + id + " остановлен")
	}

	return err, result
}
