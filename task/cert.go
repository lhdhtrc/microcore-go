package task

import (
	"fmt"
	"github.com/lhdhtrc/func-go/file"
	taskCore "github.com/lhdhtrc/task-go/core"
	taskModel "github.com/lhdhtrc/task-go/model"
	"path/filepath"
	"reflect"
)

func GetRemoteCert(task *taskCore.TaskCoreEntity, dir string, config interface{}) {
	dirPath := filepath.Join("dep", "cert", dir)

	// 遍历config的字段
	valueOfConfig := reflect.ValueOf(config).Elem()
	typeOfConfig := valueOfConfig.Type()
	for i := 0; i < valueOfConfig.NumField(); i++ {
		fieldValue := valueOfConfig.Field(i)
		fieldType := typeOfConfig.Field(i)
		if fieldValue.IsValid() && !fieldValue.IsZero() && fieldType.Type.Kind() == reflect.String {
			remote := fieldValue.String()

			// 分割路径，得到文件名部分
			f := filepath.Base(remote)
			local := filepath.Join(dirPath, f)
			fieldValue.SetString(local)

			task.Add(taskModel.TaskEntity{
				Id: fmt.Sprintf("GetRemoteCert_%d", i),
				Handle: func() error {
					read, err := file.ReadRemote(remote)
					if err != nil {
						return err
					}

					err = file.WriteLocal(local, read)
					return err
				},
			})
		}
	}
}
