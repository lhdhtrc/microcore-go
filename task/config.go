package task

import (
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/func-go/file"
	taskCore "github.com/lhdhtrc/task-go/core"
	taskModel "github.com/lhdhtrc/task-go/model"
)

func ReadRemoteConfig(task *taskCore.TaskCoreEntity, source []string, config []interface{}) {
	for i, it := range source {
		task.Add(taskModel.TaskEntity{
			Id: fmt.Sprintf("ReadRemoteConfig_%d", i),
			Handle: func() error {
				bytes, err := file.ReadRemote(it)
				if err != nil {
					return err
				}
				err = json.Unmarshal(bytes, config[i])
				return err
			},
		})
	}
}
