package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод удаляет контейнер */
func Remove(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	if err = cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{}); err == nil {
		result, err = json.Marshal("Контейнер " + id + " удален")
	}

	return err, result
}
