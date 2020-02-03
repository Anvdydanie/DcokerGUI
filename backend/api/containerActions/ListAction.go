package containerActions

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/url"
	"strconv"
)

/** Метод выводит список контейнеров */
func List(cli client.APIClient, params url.Values) (err error, result []byte) {
	var showAllParam = params.Get("showAll")
	if showAllParam != "1" {
		showAllParam = "0"
	}
	showAll, _ := strconv.ParseBool(showAllParam)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: showAll})
	if err == nil && len(containers) == 0 {
		result, err = json.Marshal("Нет ни одного запущенного контейнера")
	} else if err == nil && len(containers) > 0 {
		var _result [][]string
		for _, container := range containers {
			_result = append(_result, []string{container.ID, container.Names[0]})
		}
		result, err = json.Marshal(_result)
	}

	return err, result
}
