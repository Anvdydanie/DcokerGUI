package containerActions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/* Метод останавливает все контейнеры */
func StopAll(cli client.APIClient) (err error, result []byte) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err == nil && len(containers) == 0 {
		result, err = json.Marshal("Запущенные контейнеры не найдены")
	} else if err == nil && len(containers) > 0 {
		for _, container := range containers {
			fmt.Print("Останавливаем контейнер ", container.ID[:10], "... ")
			err = cli.ContainerStop(context.Background(), container.ID, nil)
			if err != nil {
				result, err = json.Marshal("Не удалось остановить контейнер " + container.ID[:10])
			} else {
				fmt.Print("Остановлен")
			}
		}
		if err == nil {
			result, err = json.Marshal("Все контейнеры остановлены")
		}
	}

	return err, result
}
