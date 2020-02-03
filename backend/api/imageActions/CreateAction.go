package imageActions

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"net/url"
	"os"
)

/** Метод создает образ */
func Create(cli client.APIClient, params url.Values) (err error, result []byte) {
	var image = params.Get("image")
	if len(image) == 0 {
		return errors.New("не передан параметр: image"), nil
	}

	_result, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err == nil {
		defer _result.Close()
		io.Copy(os.Stdout, _result)
		result, err = json.Marshal(_result)
	}

	return err, result
}
