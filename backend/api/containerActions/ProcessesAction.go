package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод выводит список процессов внутри контейнера */
func Processes(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	containerProcesses, err := cli.ContainerTop(context.Background(), id, []string{"aux"})
	if err == nil {
		result, err = json.Marshal(containerProcesses)
	}

	return err, result
}
