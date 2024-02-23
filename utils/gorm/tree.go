package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TreeOptions struct {
	DB      *gorm.DB
	Table   string
	AppId   string
	Id      string
	Preload []string
}

func Tree[T interface{}](options TreeOptions) (T, []*T, error) {
	parentId := options.Id

	if options.Id == "" {
		options.Id = "IS NULL"
	} else {
		options.Id = fmt.Sprintf("= '%s'", options.Id)
	}

	statement := fmt.Sprintf(`WITH RECURSIVE tree AS (SELECT * FROM %s WHERE app_id = '%s' AND parent_id %s UNION ALL SELECT t.* FROM %s t INNER JOIN tree ON tree.id = t.parent_id) SELECT * FROM tree;`, options.Table, options.AppId, options.Id, options.Table)

	var sql *gorm.DB
	var row T
	var list []*T
	sql = options.DB.Raw(statement)
	if options.Id != "IS NULL" {
		if res := options.DB.Where("id = ?", parentId).First(&row); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return row, list, res.Error
		}
	}
	for _, field := range options.Preload {
		sql.Preload(field)
	}
	sql.Find(&list)

	return row, list, nil
}
