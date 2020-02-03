package imageActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/url"
)

/** Метод инспектирует образ */
func Remove(cli client.APIClient, params url.Values) (err error, result []byte) {
	var imageName = params.Get("imageName")
	if len(imageName) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	containerInfo, err := cli.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{})
	if err == nil {
		result, err = json.Marshal(containerInfo)
	}

	return err, result
}
