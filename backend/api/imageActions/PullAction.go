package imageActions

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"net/url"
	"os"
)

/** Метод делает пул образа */
func Pull(cli client.APIClient, params url.Values) (err error, result []byte) {

	var username = params.Get("username")
	var password = params.Get("password")
	var imageName = params.Get("imageName")
	if len(username) == 0 || len(password) == 0 || len(imageName) == 0 {
		return errors.New("не передан параметр: username или password"), nil
	}

	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)

	if err == nil {
		authStr := base64.URLEncoding.EncodeToString(encodedJSON)
		imagePull, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{RegistryAuth: authStr})
		defer imagePull.Close()
		if err == nil {
			io.Copy(os.Stdout, imagePull)
			result, err = json.Marshal(imagePull)
		}
	}

	return err, result
}
