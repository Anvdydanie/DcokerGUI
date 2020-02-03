package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/client"
	"net/url"
)

/** Метод возвращает информацию о контейнере */
func Inspect(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	containerInfo, err := cli.ContainerInspect(context.Background(), id)
	if err == nil {
		result, err = json.Marshal(containerInfo)
	}

	return err, result
}
