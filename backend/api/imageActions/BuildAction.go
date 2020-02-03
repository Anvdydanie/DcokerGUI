package imageActions

import (
	"github.com/docker/docker/client"
	"net/url"
)

/** Метод создает образ */
func Build(cli client.APIClient, params url.Values) (err error, result []byte) {

	return err, result
}
