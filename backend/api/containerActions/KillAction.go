package containerActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/client"
	"net/url"
)

/* Метод убивает контейнер */
func Kill(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	if err = cli.ContainerKill(context.Background(), id, "SIGINT"); err == nil {
		result, err = json.Marshal("Контейнер " + id + " убит")
	}

	return err, result
}
