package containerActions

import (
	"github.com/docker/docker/client"
	"net/url"
)

/** Метод создает новый контейнер */
func Create(cli client.APIClient, params url.Values) (err error, result []byte) {

	/*out, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		log.Fatal(err)
	} else {
		io.Copy(os.Stdout, out)

		_result, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image: imageName,
		}, nil, nil, "")
		if err != nil {
			log.Fatal(err)
		}
		result, err = json.Marshal(_result)
		if err != nil {
			log.Fatal(err)
		}
	}*/

	return err, result
}
