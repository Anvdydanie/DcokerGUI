package imageActions

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/url"
	"strconv"
)

/** Метод выводит список образов */
func List(cli client.APIClient, params url.Values) (err error, result []byte) {
	var showAllParam = params.Get("showAll")
	if showAllParam != "1" {
		showAllParam = "0"
	}
	showAll, _ := strconv.ParseBool(showAllParam)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: showAll})
	if err == nil && len(images) == 0 {
		result, err = json.Marshal("Нет ни одного запущенного контейнера")
	} else if err == nil && len(images) > 0 {
		var _result [][]string
		for _, image := range images {
			_result = append(_result, []string{image.ID, image.RepoTags[0]})
		}
		result, err = json.Marshal(_result)
	}

	return err, result
}
