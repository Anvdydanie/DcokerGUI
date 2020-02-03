package containerActions

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

func Logs(cli client.APIClient, params url.Values) (err error, result []byte) {
	var id = params.Get("id")
	if len(id) == 0 {
		return errors.New("не передан параметр: id контейнера"), nil
	}

	var options = types.ContainerLogsOptions{ShowStdout: true}
	containerLogs, err := cli.ContainerLogs(context.Background(), id, options)
	io.Copy(os.Stdout, containerLogs)

	if err == nil {
		result, err = json.Marshal(containerLogs)
	}

	return err, result
}
