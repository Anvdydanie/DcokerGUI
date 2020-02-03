package containerActions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

/* Метод запускает контейнер */
func StartAll(cli client.APIClient) (err error, result []byte) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err == nil && len(containers) == 0 {
		if result, err = json.Marshal("Контейнеры для запуска не найдены"); err != nil {
			log.Fatal(err)
		}
	} else if err == nil && len(containers) > 0 {
		for _, container := range containers {
			fmt.Print("Запускаем контейнер ", container.ID[:10], "... ")
			err = cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
			if err == nil {
				fmt.Print("Запущен")
			}
		}
		if err == nil {
			result, err = json.Marshal("Все контейнеры успешно запущены")
		}
	}

	return err, result
}
